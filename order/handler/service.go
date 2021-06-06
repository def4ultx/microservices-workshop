package handler

import (
	"log"
	"order/inventory"
	"order/notification"
	"order/payment"
	"order/shipping"
)

type OrderService struct {
	InventoryClient    *inventory.Client
	PaymentClient      *payment.Client
	ShippingClient     *shipping.Client
	NotificationClient *notification.Client
	OrderRepository    *OrderRepository
}

type CreateOrder struct {
	CartID     int
	UserID     int
	CreditCard payment.CreditCard
}

func (s *OrderService) CreateOrder(order CreateOrder) (string, error) {
	products, err := s.InventoryClient.GetCartProducts(order.CartID)
	if err != nil {
		return "", err
	}
	totalAmount := inventory.GetTotalAmount(products)

	paymentID, err := s.PaymentClient.ChargeCreditCard(totalAmount, order.CreditCard)
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
	orderId, err := s.OrderRepository.InsertOrder(o)
	if err != nil {
		return "", err
	}

	_ = s.NotificationClient.SendEmail("test@futureskill.co", "Hello from workshop.")

	return orderId, nil

}

func (s *OrderService) GetOrderDetailByID(id string) (*OrderDetail, error) {
	order, err := s.OrderRepository.GetOrderByID(id)
	if err != nil {
		log.Println("1", err)
		return nil, err
	}

	payment, err := s.PaymentClient.GetPaymentDetail(order.PaymentID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	shipping, err := s.ShippingClient.GetShippingInfo(id)
	if err != nil {
		return nil, err
	}
	resp := OrderDetail{
		OrderID:     order.OrderID,
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		Products:    order.Products,
		Payment:     *payment,
	}
	resp.Shipping.Address = shipping.Address
	resp.Shipping.Status = shipping.Status
	return &resp, nil
}

func (s *OrderService) GetUserOrders(id int) ([]Order, error) {
	orders, err := s.OrderRepository.GetOrdersByUserID(id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
