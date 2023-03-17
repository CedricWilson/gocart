package service

import (
	"errors"
	"gocart/models"
	"gocart/utils"
	"time"

	"github.com/google/uuid"
)

func PlaceOrderService(userId string) (models.Order, error) {

	//* Getting cart items of the user
	itemList := UserCart[userId]
	delete(UserCart, userId)

	//* Deducting placed items
	for _, e := range itemList {
		v := Products[e.ID]
		v.AvailableQuantity -= e.Quantity
		Products[e.ID] = v
	}

	order := CreateOrder(itemList, userId)

	CheckForOffers(&order)

	UserOrders = append(UserOrders, order)

	return order, nil
}

func CreateOrder(itemList []models.CartProduct, userID string) models.Order {

	order := models.Order{}

	//
	order.ID = uuid.New().String()
	order.OrderStatus = utils.OrderStatus["PLACED"]
	//
	price := 0.00
	for _, e := range itemList {
		mod := models.OrderProduct{}
		mod.Name = Products[e.ID].Name
		mod.Type = Products[e.ID].Type
		mod.Quantity = e.Quantity
		price += Products[e.ID].Price
		order.Products = append(order.Products, mod)
	}
	order.OrderBaseValue = price
	order.OrderDiscountedValue = price
	//
	order.UserID = userID

	return order

}

func CheckForOffers(order *models.Order) {
	TriplePremiumProductDiscount(order)

}

func TriplePremiumProductDiscount(order *models.Order) {
	count := 0

	for _, e := range order.Products {
		if e.Type == utils.ProductType["PREMIUM"] {
			count++
		}
	}

	if count >= 3 {
		order.OrderDiscountedValue = order.OrderBaseValue - order.OrderBaseValue*(10.00/100)
	}
}

func FetchOrdersService(userId string) ([]models.Order, error) {
	orderItems := []models.Order{}

	for _, e := range UserOrders {
		if e.UserID == userId {
			orderItems = append(orderItems, e)
		}
	}

	return orderItems, nil
}

func UpdateOrderStatusService(userId, orderID string, orderStatus int) error {

	for i, e := range UserOrders {
		if e.UserID == userId && e.ID == orderID {
			
			e.OrderStatus = orderStatus
			if orderStatus == utils.OrderStatus["DISPATCHED"] {
				t := time.Now()
				e.DispatchDate = &t
			}
			UserOrders[i] = e
			return nil
		}
	}

	return errors.New("no such order found")
}
