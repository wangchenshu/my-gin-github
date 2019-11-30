package product

import (
	"mygin/httpd/common"
	"mygin/httpd/db"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productPostRequest struct {
	ProductTypeID int    `json:"product_type_id"`
	Name          string `json:"name"`
}

// Create - create a new product
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer common.MyRecover()
		_ = c.MustGet("claims").(*jwtauth.CustomClaims)

		requestBody := productPostRequest{}
		c.Bind(&requestBody)

		product := model.Products{
			ProductTypeID: requestBody.ProductTypeID,
			Name:          requestBody.Name,
		}

		err := db.Db.Create(&product)

		if product.ID == 0 {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
