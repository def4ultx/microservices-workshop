package handler

import (
	"encoding/json"
	"net/http"
	"order/api"
	"order/inventory"
	"order/notification"
	"order/payment"
	"order/shipping"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderHandler struct {
	svc *OrderService
}

func NewOrderHandler(i *inventory.Client, p *payment.Client, s *shipping.Client, n *notification.Client, db *mongo.Client) *OrderHandler {
	repo := NewOrderRepository(db)
	svc := &OrderService{i, p, s, n, repo}
	return &OrderHandler{svc: svc}
}

type CreateOrderRequest struct {
	CartID  int `json:"cartId"`
	UserID  int `json:"userId"`
	Payment struct {
		Method     string             `json:"method"`
		CreditCard payment.CreditCard `json:"creditCard"`
	} `json:"payment"`
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusBadRequest)
		return
	}

	order := CreateOrder{
		CartID:     req.CartID,
		UserID:     req.UserID,
		CreditCard: req.Payment.CreditCard,
	}
	orderID, err := h.svc.CreateOrder(order)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusInternalServerError)
		return
	}

	resp := struct {
		OrderID string `json:"orderId"`
		Status  string `json:"status"`
	}{
		OrderID: orderID,
		Status:  "Created",
	}
	api.WriteSuccessResponse(w, &resp)
}

func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		api.WriteErrorResponse(w, http.StatusBadRequest)
		return
	}

	detail, err := h.svc.GetOrderDetailByID(id)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusInternalServerError)
		return
	}

	api.WriteSuccessResponse(w, detail)
}

func (h *OrderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["userId"]
	if !ok {
		api.WriteErrorResponse(w, http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusBadRequest)
		return
	}

	orders, err := h.svc.GetUserOrders(userId)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusInternalServerError)
		return
	}

	api.WriteSuccessResponse(w, orders)
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}
