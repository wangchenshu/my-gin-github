package jwtauth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuth - jwt auth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)

		if err != nil {
			if err == ErrTokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.Header("Authorization", "Bear "+token)
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": 1, "message": err.Error()})
			return
		}

		c.Set("claims", claims)
	}
}

// JWT - jwt struct
type JWT struct {
	SigningKey []byte
}

var (
	// ErrTokenExpired - token expired
	ErrTokenExpired error = errors.New("Token is expired")
	// ErrTokenNotValidYet - token not valid yet
	ErrTokenNotValidYet error = errors.New("Token not active yet")
	// ErrTokenMalformed - token malformed
	ErrTokenMalformed error = errors.New("That's not even a token")
	// ErrTokenInvalid - token invalid
	ErrTokenInvalid error = errors.New("Couldn't handle this token")
	// SignKey - sign key
	SignKey string = "mygin"
)

// CustomClaims - custom claims
type CustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// NewJWT - new jwt
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey - get sign key
func GetSignKey() string {
	return SignKey
}

// SetSignKey - set sign key
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken - create token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken - parse token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

// RefreshToken - refresh token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}

	return "", ErrTokenInvalid
}
