package models

type CartAddModel struct {
	Products []CartProduct `binding:"required"`
}

type CartProduct struct {
	ID       int
	Quantity int
}
