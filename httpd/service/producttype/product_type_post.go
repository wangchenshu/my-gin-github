package producttype

import (
	"mygin/httpd/common"
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productTypePostRequest struct {
	Name string `json:"name"`
}

// Create - create a new product type
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		requestBody := productTypePostRequest{}
		c.Bind(&requestBody)

		productType := model.ProductType{
			Name: requestBody.Name,
		}

		db.Db.Create(&productType)

		c.JSON(http.StatusOK, productType)
	}
}
