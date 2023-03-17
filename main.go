package main

import (
	"gocart/routes"
	"gocart/service"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func init() {
	service.FillProducts()
}

func main() {
	shop := router.Group("/shop")

	//
	routes.ProductRoute(shop)
	routes.CartRoute(shop)
	routes.OrderRoute(shop)
	


	router.Run(":8080")
}
