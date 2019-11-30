package chatfuel

import (
	"fmt"
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type checkoutCartPostRequest struct {
	MessengerUserID string `json:"messenger user id"`
}

// AddCart - add a product to cart
func AddCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.ChatfuelCarts{}
		c.Bind(&requestBody)

		cart := model.Carts{
			MessengerUserID: requestBody.MessengerUserID,
			FirstName:       requestBody.FirstName,
			ProductID:       requestBody.ProductID,
			ProductName:     requestBody.ProductName,
			Qty:             requestBody.Qty,
			Price:           requestBody.Price,
		}

		db.Db.Create(&cart)

		text := []model.Text{}
		text = append(text, model.Text{
			Text: "加入購物車成功",
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}

// CheckoutCart - checkout
func CheckoutCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := clearCartPostRequest{}
		c.Bind(&requestBody)
		messengerUserID := requestBody.MessengerUserID

		fbUser := model.FbUsers{}
		db.Db.Where("messenger_user_id = ?", messengerUserID).Find(&fbUser)

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
		detail := ""
		for _, cart := range carts {
			addText := fmt.Sprintf("品名: %v, 數量: %d NT$: %d", cart.ProductName, cart.Qty, cart.Price)
			totalPrice += cart.Price
			detail += cart.ProductName + ", "

			text = append(text, model.Text{
				Text: addText,
			})
		}

		text = append(text, model.Text{
			Text: fmt.Sprintf("合計 NT$: %d", totalPrice),
		})

		order := model.Orders{
			MessengerUserID: fbUser.MessengerUserID,
			FirstName:       fbUser.FirstName,
			Detail:          detail,
			TotalPrice:      totalPrice,
		}

		// create order
		db.Db.Create(&order)

		// clear carts
		db.Db.Where("messenger_user_id = ?", messengerUserID).Delete(model.Carts{})

		text = append(text, model.Text{
			Text: "結帳完成，感謝您的購買。",
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}
