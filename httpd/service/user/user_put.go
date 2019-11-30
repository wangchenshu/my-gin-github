package user

import (
	"mygin/httpd/common"
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userPutRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Update - update user info by id
func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		requestBody := userPutRequest{}
		c.Bind(&requestBody)

		user := model.Users{}
		db.Db.Where("id = ?", requestBody.ID).First(&user)

		if user.ID != 0 {
			hash, _ := common.HashPassword(requestBody.Password)
			user.Name = requestBody.Name
			user.Password = hash

			db.Db.Save(&user)

			c.JSON(http.StatusOK, user)
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "user not found"})
	}
}
