package controller

import (
	"gocart/models"
	"gocart/service"
	"gocart/utils"


	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var request models.CartAddModel
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	//* User is predefined
	userID := utils.UserID
	// 
	err = service.AddToCartService(request.Products,userID)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "SUCCESS")
}

func FetchCart(c *gin.Context) {
	//* User is predefined
	userId:= utils.UserID
	items, err := service.FetchCartService(userId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, items)
}
