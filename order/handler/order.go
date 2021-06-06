package handler

import (
	"order/inventory"
	"order/payment"
)

type Order struct {
	OrderID     string
	Status      string
	UserID      int
	PaymentID   int
	TotalAmount int
	Products    []inventory.Product
}

type OrderDetail struct {
	OrderID     string                `json:"orderId"`
	Status      string                `json:"status"`
	TotalAmount int                   `json:"totalAmount"`
	Products    []inventory.Product   `json:"products"`
	Payment     payment.PaymentDetail `json:"payment"`
	Shipping    struct {
		Address string `json:"address"`
		Status  string `json:"status"`
	} `json:"shipping"`
}
