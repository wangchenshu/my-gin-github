package user

import (
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"mygin/httpd/common"
)

// Get - get user by id
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		id := c.Param("id")
		user := model.Users{}
		db.Db.Where("id = ?", id).Find(&user)

		if user.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"ID": ""})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// GetAll - get all users
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		users := []model.Users{}
		db.Db.Find(&users)

		c.JSON(http.StatusOK, users)

		c.JSON(http.StatusOK, gin.H{"data": nil})
	}
}
