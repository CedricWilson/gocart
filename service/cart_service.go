package service

import (
	"gocart/models"

)

func AddToCartService(products []models.CartProduct,userID string) error {
	UserCart[userID] = products
	return nil
}
func FetchCartService(userID string) ([]models.CartProduct, error) {
	cartItems := []models.CartProduct{}
	cartItems = append(cartItems, UserCart[userID]...)

	return cartItems, nil
}
