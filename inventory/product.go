package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductHandler struct {
	db *pgxpool.Pool
}

type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

func (h *ProductHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) ListProduct(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Products []Product `json:"products"`
	}

	stmt := `SELECT id, name, price, amount FROM products`
	rows, err := h.db.Query(context.Background(), stmt)
	if err != nil {
		log.Println("got err: ", err)
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := make([]Product, 0)
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Amount)
		if err != nil {
			break
		}

		products = append(products, p)
	}

	err = rows.Err()
	if err != nil {
		log.Println("got err: ", err)
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	resp := Response{products}
	writeSuccessResponse(w, &resp)
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	var id int64
	stmt := `INSERT INTO products(name, price, amount) VALUES ($1, $2, $3) RETURNING id`
	err = h.db.QueryRow(context.Background(), stmt, p.Name, p.Price, p.Amount).Scan(&id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	p.ID = id
	writeSuccessResponse(w, &p)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	var p Product
	stmt := `SELECT id, name, price, amount FROM products WHERE id = $1`
	err = h.db.QueryRow(context.Background(), stmt, &id).Scan(&p.ID, &p.Name, &p.Price, &p.Amount)
	if errors.Is(err, pgx.ErrNoRows) {
		writeErrorResponse(w, http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("got err: ", err)
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}

	writeSuccessResponse(w, &p)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	var p Product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	p.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	stmt := `UPDATE products SET (name, price, amount) = ($2, $3, $4) WHERE id = $1`
	res, err := h.db.Exec(context.Background(), stmt, p.ID, p.Name, p.Price, p.Amount)
	if err != nil {
		log.Println("got err: ", err)
		writeErrorResponse(w, http.StatusInternalServerError)
		return
	}
	if res.RowsAffected() == 0 {
		writeErrorResponse(w, http.StatusNotFound)
		return
	}

	writeSuccessResponse(w, &p)
}
