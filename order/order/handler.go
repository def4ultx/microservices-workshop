package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/api"
	"order/payment"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc Service
}

type CreateOrderRequest struct {
	CartID  int                 `json:"cartId"`
	UserID  int                 `json:"userId"`
	Payment OrderPaymentRequest `json:"payment"`
}

type OrderPaymentRequest struct {
	Method     string             `json:"method"`
	CreditCard payment.CreditCard `json:"creditCard"`
}

func (h *Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		api.WriteErrorResponse(w, http.StatusBadRequest)
		return
	}

	order, err := h.svc.GetOrderDetailByID(id)
	if err != nil {
		api.WriteErrorResponse(w, http.StatusInternalServerError)
		return
	}

	api.WriteSuccessResponse(w, order)
}

func (h *Handler) ListOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}
