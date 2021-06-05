package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/api"
	"order/inventory"
	"order/payment"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	svc Service
}

func NewHandler(i *inventory.Client, pc *payment.Client, db *mongo.Client) *Handler {
	repo := NewRepository(db)
	svc := &service{i, pc, repo}
	return &Handler{svc}
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

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented")
}
