package model

import "github.com/midtrans/midtrans-go/coreapi"

// import "github.com/midtrans/midtrans-go/coreapi"

type PaymentRequest struct {
	OrderID string
	Amount  int
}

type PaymentResponse struct {
	Status        []coreapi.VANumber
	TransactionID string
}
