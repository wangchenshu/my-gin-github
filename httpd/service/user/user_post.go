package user

import (
	"fmt"
	"mygin/httpd/common"
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwtauth "mygin/httpd/middleware/jwt"
)

type userPostRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var hmacSampleSecret []byte

// Create - create a new user
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		requestBody := userPostRequest{}
		c.Bind(&requestBody)

		hash, _ := common.HashPassword(requestBody.Password)

		user := model.Users{
			Name:     requestBody.Name,
			Password: hash,
		}

		db.Db.Create(&user)

		c.JSON(http.StatusOK, user)
	}
}

// Login - user login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := userPostRequest{}
		c.Bind(&requestBody)

		user := model.Users{}
		db.Db.Where("name = ?", requestBody.Name).First(&user)

		if user.ID != 0 {
			match := common.CheckPasswordHash(requestBody.Password, user.Password)

			if match == true {
				j := &jwtauth.JWT{
					SigningKey: []byte(jwtauth.GetSignKey()),
				}

				claims := jwtauth.CustomClaims{
					Name:     requestBody.Name,
					Password: requestBody.Password,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
						Issuer:    jwtauth.GetSignKey(),
					},
				}

				token, err := j.CreateToken(claims)
				if err != nil {
					c.String(http.StatusOK, err.Error())
					c.Abort()
				}

				c.String(http.StatusOK, token+"---------------<br>")
				res, err := j.ParseToken(token)

				if err != nil {
					if err == jwtauth.ErrTokenExpired {
						newToken, err := j.RefreshToken(token)
						if err != nil {
							c.String(http.StatusOK, err.Error())
						} else {
							c.String(http.StatusOK, newToken)
						}
					} else {
						c.String(http.StatusOK, err.Error())
					}
				} else {
					c.JSON(http.StatusOK, res)
					return
				}
			} else {
				c.JSON(http.StatusOK, gin.H{"token": nil})
				return
			}
		}
	}
}

// LoginTest - user login test
func LoginTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := userPostRequest{}
		c.Bind(&requestBody)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": requestBody.Name,
			"nbf":      time.Date(2019, 11, 30, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)

		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

// ValidateToken - validate token
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
		tokenString := token

		ok, err := ValidateTokenMiddleware(tokenString)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"status": err})
		}

		c.JSON(http.StatusOK, gin.H{"status": ok})
	}
}

// ValidateTokenMiddleware - validate token
func ValidateTokenMiddleware(tokenString string) (bool, error) {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["nbf"])
		return true, nil
	} else {
		fmt.Println(err)
		return false, err
	}
}
