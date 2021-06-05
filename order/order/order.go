package order

import (
	"order/inventory"
)

type Order struct {
	ID          string
	UserID      int
	Status      string
	TotalAmount int
	Products    []inventory.Product
	PaymentID   int
}
