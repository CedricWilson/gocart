package models

import "time"

type Order struct {
	ID                   string
	UserID               string
	OrderDiscountedValue float64
	OrderBaseValue       float64
	Products             []OrderProduct
	OrderStatus          int
	DispatchDate         *time.Time
}

type OrderProduct struct {
	Name        string
	PlacedPrice float64
	Quantity    int
	Type        int
}
