package chatfuel

import (
	"mygin/httpd/db"
	"mygin/httpd/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FbUserCreate - create a new user
func FbUserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := model.ChatfuelFbUser{}
		c.Bind(&requestBody)

		fbUser := model.FbUsers{
			MessengerUserID: requestBody.MessengerUserID,
			FirstName:       requestBody.FirstName,
			LastName:        requestBody.LastName,
			Gender:          requestBody.Gender,
			ProfilePicURL:   requestBody.ProfilePicURL,
			Timezone:        requestBody.Timezone,
			Locale:          requestBody.Locale,
		}

		db.Db.Create(&fbUser)

		text := []model.Text{}
		text = append(text, model.Text{
			Text: "加入成功",
		})

		message := model.Message{
			Message: text,
		}

		c.JSON(http.StatusOK, message)
	}
}
