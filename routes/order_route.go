package routes

import (
	"github.com/gin-gonic/gin"
	"gocart/controllers"
)

func OrderRoute(router *gin.RouterGroup) {
	api := router.Group("/orders")
	//
	api.POST("/place", controller.PlaceOrder)
	api.GET("/fetch", controller.FetchOrders)
	api.POST("/update", controller.UpdateOrders)

}
