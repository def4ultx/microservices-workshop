package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CartHandler struct {
	db *pgxpool.Pool
}

func (h *CartHandler) Create(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		UserID int64 `json:"userId"`
	}

	type Response struct {
		CartID int64 `json:"cartId"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	var resp Response
	stmt := `INSERT INTO carts(user_id) VALUES ($1) RETURNING id`
	err = h.db.QueryRow(context.Background(), stmt, req.UserID).Scan(&resp.CartID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, &resp)
}

func (h *CartHandler) Get(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		Products []Product `json:"products"`
	}

	cartId, ok := mux.Vars(r)["cartId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	stmt := `SELECT c.product_id, c.amount, p.name, p.price
				FROM cart_products AS c
				LEFT JOIN products AS p
				ON c.product_id = p.id
				WHERE c.cart_id = $1`
	rows, err := h.db.Query(context.Background(), stmt, cartId)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var p Product

		err = rows.Scan(&p.ID, &p.Amount, &p.Name, &p.Price)
		if err != nil {
			break
		}

		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	resp := Response{products}
	writeSuccessResponse(w, &resp)
}

func (h *CartHandler) RemoveAllProduct(w http.ResponseWriter, r *http.Request) {

	cartId, ok := mux.Vars(r)["cartId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	stmt := `DELETE FROM cart_products WHERE cart_id = $1`
	_, err := h.db.Exec(context.Background(), stmt, cartId)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, struct{}{})
}

func (h *CartHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ProductID int64 `json:"productId"`
		Amount    int   `json:"amount"`
	}

	cartId, ok := mux.Vars(r)["cartId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	productId, ok := mux.Vars(r)["productId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	stmt := `INSERT INTO cart_products (cart_id, product_id, amount) VALUES ($1, $2, $3)
				ON CONFLICT (cart_id, product_id) DO UPDATE SET amount = excluded.amount`
	_, err = h.db.Exec(context.Background(), stmt, cartId, productId, req.Amount)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, struct{}{})
}

func (h *CartHandler) RemoveProduct(w http.ResponseWriter, r *http.Request) {

	cartId, ok := mux.Vars(r)["cartId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	productId, ok := mux.Vars(r)["productId"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	stmt := `DELETE FROM cart_products WHERE cart_id = $1 AND product_id = $2`
	_, err := h.db.Exec(context.Background(), stmt, cartId, productId)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, struct{}{})
}
