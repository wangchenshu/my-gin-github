package chatfuel

import (
	"github.com/gin-gonic/gin"
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"
)

// Get - get product by id
func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product := model.Products{}
		db.Db.Where("id = ?", id).Find(&product)

		if product.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"ID": ""})
			return
		}

		text := []model.Text{}
		text = append(text, model.Text{
			Text: product.Name,
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}

// GetAll - get all products
func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := []model.Products{}
		db.Db.Find(&products)

		text := []model.Text{}

		for _, product := range products {
			text = append(text, model.Text{
				Text: product.Name,
			})
		}

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}
