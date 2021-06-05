package order

import (
	"fmt"
	"net/http"
	"order/api"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc Service
}

func (h *Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// call inventory
	// call to payment
	// save to db
	// produce to noti and shipping
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
