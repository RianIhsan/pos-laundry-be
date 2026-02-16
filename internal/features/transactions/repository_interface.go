package transactions

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context, transaction entities.Transaction, items []entities.TransactionItem) (entities.Transaction, error)
	GetList(ctx context.Context) ([]entities.Transaction, error)
	FindById(ctx context.Context, transactionId uint) (entities.Transaction, error)
	Update(ctx context.Context, transactionId uint, data entities.Transaction) (entities.Transaction, error)
}
