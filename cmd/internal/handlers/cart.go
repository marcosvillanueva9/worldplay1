package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/models"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/services"
	"net/http"
	"strconv"
)

func CreateCart(c *gin.Context) {
	strCartId, ok := c.Params.Get("cart_id")
	if !ok {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	cartId, err := strconv.Atoi(strCartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id isn´t a number")
		return
	}

	err = services.CreateCart(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
	return
}

func AddProductsToCart(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	strCartId, ok := c.Params.Get("cart_id")
	if !ok {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	cartId, err := strconv.Atoi(strCartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id isn´t a number")
		return
	}

	err = services.AddProductToCart(product, cartId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetCart(c *gin.Context) {
	strCartId, ok := c.Params.Get("cart_id")
	if !ok {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	cartId, err := strconv.Atoi(strCartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id isn´t a number")
		return
	}

	cart, err := services.GetCart(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, cart)
	return
}
