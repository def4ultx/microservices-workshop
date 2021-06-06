package order

import (
	"log"
	"order/inventory"
	"order/payment"
)

type Service interface {
	CreateOrder(CreateOrder) (string, error)
	GetOrderDetailByID(string) (*OrderDetail, error)
	GetUserOrders(int) ([]Order, error)
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

	// push message to noti/shipping topic

	return orderId, nil

}

func (s *service) GetOrderDetailByID(id string) (*OrderDetail, error) {
	order, err := s.orderRepository.GetOrderByID(id)
	if err != nil {
		log.Println("1", err)
		return nil, err
	}

	payment, err := s.paymentClient.GetPaymentDetail(order.PaymentID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// shipping, err := s.shippingClient.GetShippingInfo(id)
	// if err != nil {
	// 	return nil, err
	// }
	resp := OrderDetail{
		OrderID:     order.OrderID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Products:    order.Products,
		Payment:     *payment,
	}
	return &resp, nil
}

func (s *service) GetUserOrders(id int) ([]Order, error) {
	orders, err := s.orderRepository.GetOrdersByUserID(id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
