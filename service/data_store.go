package service

import (
	"gocart/models"
	"gocart/utils"
)

var (
	Products = make(map[int]models.Product)
	UserCart = make(map[string][]models.CartProduct)
	UserOrders []models.Order

)

func FillProducts() {
	Products = map[int]models.Product{
		1: {
			ID:                1,
			Name:              "CHIPS",
			Price:             100,
			Type:              utils.ProductType["REGULAR"],
			AvailableQuantity: 10,
		},
		2: {
			ID: 2,
			Name:              "Water",
			Price:             100,
			Type:              utils.ProductType["BUDGET"],
			AvailableQuantity: 10,
		},
		3: {
			ID:                3,
			Name:              "Cheese",
			Price:             100,
			Type:              utils.ProductType["PREMIUM"],
			AvailableQuantity: 10,
		},
		4: {
			ID:                4,
			Name:              "TOAST",
			Price:             100,
			AvailableQuantity: 10,
			Type:              utils.ProductType["PREMIUM"],
		},
		5: {
			ID:                5,
			Name:              "BURGER",
			Price:             100,
			AvailableQuantity: 10,
			Type:              utils.ProductType["PREMIUM"],
		},
		6: {
			ID:                6,
			Name:              "CAKE",
			Price:             100,
			AvailableQuantity: 10,
			Type:              utils.ProductType["PREMIUM"],
		},
	}

}
