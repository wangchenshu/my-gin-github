package routes

import (
	"github.com/gin-gonic/gin"
	jwtauth "mygin/httpd/middleware/jwt"
	"mygin/httpd/service/chatfuel"
	"mygin/httpd/service/product"
	"mygin/httpd/service/producttype"
	"mygin/httpd/service/user"
)

// Engine - engine
func Engine() *gin.Engine {
	r := gin.Default()

	// Chatfuel
	r.GET("/api/chatfuel/products", chatfuel.GetAll())
	r.GET("/api/chatfuel/product/:id", chatfuel.Get())
	r.GET("/api/chatfuel/cart/:messenger_user_id", chatfuel.GetCarts())
	r.POST("/api/chatfuel/cart", chatfuel.AddCart())
	r.POST("/api/chatfuel/cart/checkout", chatfuel.CheckoutCart())

	// chatfuel free version only support get and post @@
	// r.DELETE("/api/chatfuel/cart", chatfuel.ClearCart())
	r.POST("/api/chatfuel/cart/clear", chatfuel.ClearCart())
	r.POST("/api/chatfuel/fb-user", chatfuel.FbUserCreate())

	r.POST("/api/user/login", user.Login())

	authorize := r.Group("/api/", jwtauth.JWTAuth())
	{
		// User
		authorize.POST("/user", user.Create())
		authorize.PUT("/user", user.Update())
		authorize.GET("users", user.GetAll())
		authorize.GET("/user/:id", user.Get())

		// Product Type
		authorize.GET("/product-types", producttype.GetAll())
		authorize.GET("/product-type/:id", producttype.Get())
		authorize.POST("/product-type", producttype.Create())

		// Product
		authorize.GET("/products", product.GetAll())
		authorize.GET("/product/:id", product.Get())
		authorize.POST("/product", product.Create())
	}

	return r
}
