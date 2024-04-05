package main

import (
	"log"
	"net/http"
	"midtrans/config"
	"midtrans/handler"
)

func main() {
    // Memuat konfigurasi aplikasi
    conf := config.LoadConfig()

    // Konfigurasi routing
    http.HandleFunc("/payment", func(w http.ResponseWriter, r *http.Request) {
        // Handler untuk endpoint /payment
        // Anda dapat menggunakan konfigurasi atau data lain dari `conf` di sini
        handler.PaymentHandler(w, r, conf)
    })

    // Menjalankan server HTTP
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

