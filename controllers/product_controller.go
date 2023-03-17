package controller

import (
	"github.com/gin-gonic/gin"
	"gocart/service"
)

func GetProductListings(c *gin.Context) {
	listings, err := service.GetProductListingsService()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, listings)
}
