package integration

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/handlers"
	"net/http"
)

func Run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/api/cart/create/:cart_id", handlers.CreateCart)
	r.POST("/api/cart/products/:cart_id", handlers.AddProductsToCart)
	r.GET("/api/cart/:cart_id", handlers.GetCart)
	r.POST("/api/order/:order_id", handlers.CreateOrder)
	r.GET("/api/order/:order_id", handlers.GetOrderProducts)

	r.Run()
}
