package order

import (
	"order/inventory"
	"order/payment"
)

type Service interface {
	CreateOrder(CreateOrder) (string, error)
	GetOrderDetailByID(id string) (*OrderDetail, error)
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
func (s *service) GetOrderDetailByID(id string) (*OrderDetail, error) {
	// order, err := s.orderRepository.GetOrderByID(id)
	// if err != nil {
	// 	return nil, err
	// }

	// detail, err := s.paymentClient.GetPaymentDetail(order.Payment.ID)
	// if err != nil {
	// 	return nil, err
	// }

	// shipping, err := s.shippingClient.GetShippingInfo(id)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
