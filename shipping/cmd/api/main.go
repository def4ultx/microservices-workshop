package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the service")

	r := mux.NewRouter()
	r.HandleFunc("/shipping/{orderId}", GetShippingHandler).Methods(http.MethodGet)
	r.HandleFunc("/shipping/{orderId}", UpdateShippingHandler).Methods(http.MethodPut)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Println("The service is ready to listen and serve.")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type ShippingInformation struct {
	Address string `json:"address"`
	Status  string `json:"status"`
}

func GetShippingHandler(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ShippingInformation
	}

	resp := Response{
		ShippingInformation{
			Address: "111/11 Bangkok 10101",
			Status:  "Completed",
		},
	}

	writeSuccessResponse(w, &resp)
}

func UpdateShippingHandler(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ShippingInformation
	}

	type Response struct {
		ShippingInformation
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest)
		return
	}

	writeSuccessResponse(w, &req)
}

func writeSuccessResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func writeErrorResponse(w http.ResponseWriter, code int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(struct{}{})
}
