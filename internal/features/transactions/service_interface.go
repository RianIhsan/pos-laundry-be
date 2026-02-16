package transactions

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions/dto"
)

type TransactionServiceInterface interface {
	CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest, userID uint) (dto.TransactionResponse, error)
	GetList(ctx context.Context) ([]dto.TransactionResponse, error)
	GetById(ctx context.Context, transactionId uint) (dto.TransactionResponse, error)
	UpdateStatus(ctx context.Context, transactionId uint, userID uint, req dto.UpdateTransactionStatusRequest) (dto.TransactionResponse, error)
}
