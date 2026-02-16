package dto

import (
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type TransactionItemResponse struct {
	ID          uint    `json:"id"`
	ServiceID   uint    `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Qty         float64 `json:"qty"`
	PriceAtTime float64 `json:"price_at_time"`
	Note        string  `json:"note"`
	Subtotal    float64 `json:"subtotal"`
}

type TransactionResponse struct {
	ID            uint                      `json:"id"`
	InvoiceNo     string                    `json:"invoice_no"`
	CustomerID    uint                      `json:"customer_id"`
	CustomerName  string                    `json:"customer_name"`
	UserID        uint                      `json:"user_id"`
	UserName      string                    `json:"user_name"`
	TotalPrice    float64                   `json:"total_price"`
	PaymentMethod string                    `json:"payment_method"`
	PaymentStatus string                    `json:"payment_status"`
	OrderStatus   string                    `json:"order_status"`
	AmountPaid    float64                   `json:"amount_paid"`
	CreatedAt     time.Time                 `json:"created_at"`
	Items         []TransactionItemResponse `json:"items"`
}

func ToTransactionResponse(transaction entities.Transaction) TransactionResponse {
	items := make([]TransactionItemResponse, 0)

	for _, item := range transaction.Items {
		serviceName := ""
		if item.Service.ID > 0 {
			serviceName = item.Service.Name
		}
		items = append(items, TransactionItemResponse{
			ID:          item.ID,
			ServiceID:   item.ServiceID,
			ServiceName: serviceName,
			Qty:         item.Qty,
			PriceAtTime: item.PriceAtTime,
			Note:        item.Notes,
			Subtotal:    item.Subtotal,
		})
	}

	return TransactionResponse{
		ID:            transaction.ID,
		InvoiceNo:     transaction.InvoiceNo,
		CustomerID:    transaction.CustomerID,
		CustomerName:  transaction.Customer.Name,
		UserID:        transaction.UserID,
		UserName:      transaction.User.Name,
		TotalPrice:    transaction.TotalPrice,
		PaymentMethod: transaction.PaymentMethod,
		PaymentStatus: transaction.PaymentStatus,
		OrderStatus:   transaction.OrderStatus,
		AmountPaid:    transaction.AmountPaid,
		CreatedAt:     transaction.CreatedAt,
		Items:         items,
	}
}
