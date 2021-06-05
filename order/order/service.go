package order

import (
	"order/inventory"
	"order/payment"
)

type Service interface {
	CreateOrder(request CreateOrder) (string, error)
	GetOrderDetailByID(id string) (OrderDetail, error)
	ListOrderByUserID()
}

type service struct {
	inventoryClient *inventory.Client
	paymentClient   *payment.Client
	orderRepository Repository
}

type CreateOrder struct {
	CartID     int
	UserID     int
	CreditCard payment.CreditCard
}

func (s *service) CreateOrder(order CreateOrder) (string, error) {
	products, err := s.inventoryClient.GetCartProducts(order.CartID)
	if err != nil {
		return "", err
	}
	totalAmount := inventory.GetTotalAmount(products)

	paymentID, err := s.paymentClient.ChargeCreditCard(totalAmount, order.CreditCard)
	if err != nil {
		return "", err
	}

	o := Order{
		UserID:      order.UserID,
		Status:      "Created",
		TotalAmount: totalAmount,
		Products:    products,
		PaymentID:   paymentID,
	}
	orderId, err := s.orderRepository.InsertOrder(o)
	if err != nil {
		return "", err
	}

	// push message to noti topic

	return orderId, nil

}

type OrderDetail struct {
}
