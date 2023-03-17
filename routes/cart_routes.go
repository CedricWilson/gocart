package routes

import (
	"github.com/gin-gonic/gin"
	"gocart/controllers"
)

func CartRoute(router *gin.RouterGroup) {
	api := router.Group("/cart")
	//
	api.POST("/add", controller.AddToCart)
	api.GET("/fetch", controller.FetchCart)

}
