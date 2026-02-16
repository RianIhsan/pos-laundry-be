package dto

import (
	"strings"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type CreateTransactionItemRequest struct {
	ServiceID   uint    `json:"service_id" validate:"required,gt=0"`
	Qty         float64 `json:"qty" validate:"required,gt=0"`
	PriceAtTime float64 `json:"price_at_time" validate:"required,gt=0"`
	Note        string  `json:"note"`
}

type CreateTransactionRequest struct {
	CustomerID    uint                           `json:"customer_id" validate:"required,gt=0"`
	PaymentMethod string                         `json:"payment_method" validate:"required,oneof=CASH QRIS"`
	PaymentStatus string                         `json:"payment_status" validate:"required,oneof=PAID UNPAID"`
	AmountPaid    float64                        `json:"amount_paid" validate:"required,gte=0"`
	Items         []CreateTransactionItemRequest `json:"items" validate:"required,min=1,dive"`
}

func (r *CreateTransactionRequest) PrepareCreate() {
	r.PaymentMethod = strings.ToUpper(strings.TrimSpace(r.PaymentMethod))
	r.PaymentStatus = strings.ToUpper(strings.TrimSpace(r.PaymentStatus))

	for i := range r.Items {
		r.Items[i].Note = strings.TrimSpace(r.Items[i].Note)
	}
}

func (r *CreateTransactionRequest) CalculateTotalPrice() float64 {
	total := 0.0
	for _, item := range r.Items {
		total += item.Qty * item.PriceAtTime
	}
	return total
}

func ConvertToTransactionEntity(req CreateTransactionRequest, invoiceNo string, userID uint) entities.Transaction {
	totalPrice := req.CalculateTotalPrice()

	transaction := entities.Transaction{
		InvoiceNo:     invoiceNo,
		CustomerID:    req.CustomerID,
		UserID:        userID,
		TotalPrice:    totalPrice,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: req.PaymentStatus,
		AmountPaid:    req.AmountPaid,
		OrderStatus:   "WASHING", // Default status
	}

	return transaction
}

func ConvertToTransactionItemEntity(transactionID uint, req CreateTransactionItemRequest) entities.TransactionItem {
	subtotal := req.Qty * req.PriceAtTime

	return entities.TransactionItem{
		TransactionID: transactionID,
		ServiceID:     req.ServiceID,
		Qty:           req.Qty,
		PriceAtTime:   req.PriceAtTime,
		Notes:         req.Note,
		Subtotal:      subtotal,
	}
}

type UpdateTransactionStatusRequest struct {
	OrderStatus   string `json:"order_status" validate:"omitempty,oneof=WASHING READY_TO_PICKUP COMPLETED"`
	PaymentStatus string `json:"payment_status" validate:"omitempty,oneof=PAID UNPAID"`
}

func (r *UpdateTransactionStatusRequest) PrepareUpdate() {
	r.OrderStatus = strings.ToUpper(strings.TrimSpace(r.OrderStatus))
	r.PaymentStatus = strings.ToUpper(strings.TrimSpace(r.PaymentStatus))
}
