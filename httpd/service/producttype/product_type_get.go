package producttype

import (
	"mygin/httpd/common"
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get - get product type by id
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		id := c.Param("id")
		productType := model.ProductType{}
		db.Db.Where("id = ?", id).Find(&productType)

		if productType.ProductTypeID == 0 {
			c.JSON(http.StatusOK, gin.H{"ID": ""})
			return
		}

		c.JSON(http.StatusOK, productType)
	}
}

// GetAll - get all product types
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		productTypes := []model.ProductType{}
		db.Db.Find(&productTypes)

		c.JSON(http.StatusOK, productTypes)
	}
}
