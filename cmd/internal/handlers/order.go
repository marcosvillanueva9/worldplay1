package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/services"
	"net/http"
	"strconv"
)

func GetOrderProducts(c *gin.Context) {
	strCartId, ok := c.Params.Get("order_id")
	if !ok {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	cartId, err := strconv.Atoi(strCartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id isn´t a number")
		return
	}

	cart, err := services.GetOrder(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
	return
}

func CreateOrder(c *gin.Context) {
	strOrderId, ok := c.Params.Get("cart_id")
	if !ok {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	orderId, err := strconv.Atoi(strOrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id isn´t a number")
		return
	}

	err = services.CreateOrder(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
	return
}

