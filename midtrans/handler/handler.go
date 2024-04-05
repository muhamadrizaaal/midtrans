package handler

import (
	"encoding/json"
	"midtrans/config"
	"midtrans/model"
	"midtrans/services"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request, conf config.Config) {
	// Parse request body
	var paymentRequest model.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Memuat konfigurasi aplikasi
	conf = config.LoadConfig()

	// Proses pembayaran
	paymentService := services.PaymentService{
		ServerKey: conf.ServerKey,
	}
	paymentResponse, err := paymentService.ProcessPayment(paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Menyimpan response ke dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResponse)
}
