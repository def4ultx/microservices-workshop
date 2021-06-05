package order

import "order/inventory"

type Order struct {
	ID          int                 `json:"orderId"`
	Status      string              `json:"status"`
	TotalAmount int                 `json:"totalAmount"`
	Products    []inventory.Product `json:"products"`
}
