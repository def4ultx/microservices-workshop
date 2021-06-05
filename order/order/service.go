package order

type Service interface {
	GetOrderDetailByID(id string) (OrderDetail, error)
}
