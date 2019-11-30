package chatfuel

import (
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productPostRequest struct {
	ProductTypeID int    `json:"product_type_id"`
	Name          string `json:"name"`
}

type cartPostRequest struct {
	MessengerUserID string `json:"messenger user id"`
	FirstName       string `json:"first name"`
	ProductID       string `json:"product_id"`
	Qty             string `json:"qty"`
}

// Create - create a new product
func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
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
