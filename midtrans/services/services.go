package services

import (
	"log"
	"midtrans/model"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentService struct {
	ServerKey string // Menambahkan field untuk menyimpan Server Key
}

// ProcessPayment is used to process the payment
func (ps *PaymentService) ProcessPayment(request model.PaymentRequest) (model.PaymentResponse, error) {
	// Set Midtrans server key
	midtrans.ServerKey = "SB-Mid-server-wuGVT1eVnpSpoEzIY9fY_kJb"

	// Set Midtrans environment
	midtrans.Environment = midtrans.Sandbox

	// var payment model.PaymentRequest
	// Payment information
	// chargeReq := &coreapi.ChargeReq{
	// 	PaymentType: coreapi.PaymentTypeQris,
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  request.OrderID,
	// 		GrossAmt: int64(request.Amount),
	// 	},
	// }

	log.Print(request.OrderID)
	log.Print(request.Amount)
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: int64(request.Amount),
		}, BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
			// VaNumber: "1111111",
			// Bca: &coreapi.BcaBankTransferDetail{
			// 	SubCompanyCode: "0000",
			// },
		},
	}

	// Process the payment
	chargeResp, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	// Create payment response
	response := model.PaymentResponse{
		Status:        chargeResp.VaNumbers,
		TransactionID: chargeResp.TransactionID,
		// Add other fields as needed
	}

	return response, nil
}
