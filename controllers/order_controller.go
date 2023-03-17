package controller

import (
	"gocart/models"
	"gocart/service"
	"gocart/utils"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	
	//* User is predefined
	userID := utils.UserID
	//
	order, err := service.PlaceOrderService(userID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, map[string]any{
		"message": "Success",
		"details": order,
	})
}

func FetchOrders(c *gin.Context) {
	//* User is predefined
	userId := utils.UserID
	items, err := service.FetchOrdersService(userId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, items)
}

func UpdateOrders(c *gin.Context) {
	var request models.OrderUpdate

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, err)
		return
	}

	//* User is predefined
	userId := utils.UserID
	err = service.UpdateOrderStatusService(userId, request.ID, request.OrderStatus)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, "SUCCESS")
}
