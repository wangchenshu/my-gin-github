package chatfuel

import (
	"fmt"
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCarts - get cars by messenger_user_id
func GetCarts() gin.HandlerFunc {
	return func(c *gin.Context) {
		messengerUserID := c.Param("messenger_user_id")
		carts := []model.Carts{}
		db.Db.Where("messenger_user_id = ?", messengerUserID).Find(&carts)
		text := []model.Text{}

		if len(carts) < 1 {
			text = append(text, model.Text{
				Text: "購物車是空的",
			})
			message := model.Message{
				Message: text,
			}
			c.JSON(http.StatusOK, message)
			return
		}

		text = append(text, model.Text{
			Text: "項目如下: ",
		})
		totalPrice := 0
		for _, cart := range carts {
			addText := fmt.Sprintf("品名: %v, 數量: %d NT$: %d", cart.ProductName, cart.Qty, cart.Price)
			totalPrice += cart.Price
			text = append(text, model.Text{
				Text: addText,
			})
		}

		text = append(text, model.Text{
			Text: fmt.Sprintf("合計 NT$: %d", totalPrice),
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}
