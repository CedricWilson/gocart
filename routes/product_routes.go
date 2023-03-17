package routes

import (
	"gocart/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.RouterGroup) {
	api := router.Group("/product")
	//
	api.GET("/listings", controller.GetProductListings)

}
