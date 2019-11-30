package product

import (
	"mygin/httpd/common"
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get - get product by id
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		id := c.Param("id")
		product := model.Products{}
		db.Db.Where("id = ?", id).Find(&product)

		if product.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"ID": ""})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// GetAll - get all products
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		products := []model.Products{}
		db.Db.Find(&products)

		c.JSON(http.StatusOK, products)
	}
}
