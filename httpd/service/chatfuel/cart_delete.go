package chatfuel

import (
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clearCartPostRequest struct {
	MessengerUserID string `json:"messenger user id"`
}

// ClearCart - clear cart
func ClearCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := clearCartPostRequest{}
		c.Bind(&requestBody)

		db.Db.Where("messenger_user_id = ?", requestBody.MessengerUserID).Delete(model.Carts{})

		text := []model.Text{}
		text = append(text, model.Text{
			Text: "清除購物車成功",
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}
