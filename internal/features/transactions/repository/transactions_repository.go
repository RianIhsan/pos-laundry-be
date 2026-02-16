package repository

import (
	"context"
	"fmt"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type transactionPostgresRepository struct {
	db *gorm.DB
}

func NewTransactionPostgresRepository(db *gorm.DB) transactions.TransactionRepositoryInterface {
	return &transactionPostgresRepository{db: db}
}

func (r *transactionPostgresRepository) Create(ctx context.Context, transaction entities.Transaction, items []entities.TransactionItem) (entities.Transaction, error) {
	// Start transaction to ensure atomicity
	tx := r.db.WithContext(ctx).Begin()

	// Create transaction
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return entities.Transaction{}, errors.Wrap(err, "failed to create transaction")
	}

	// Create transaction items
	for i := range items {
		items[i].TransactionID = transaction.ID
		if err := tx.Create(&items[i]).Error; err != nil {
			tx.Rollback()
			return entities.Transaction{}, errors.Wrap(err, "failed to create transaction items")
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return entities.Transaction{}, errors.Wrap(err, "failed to commit transaction")
	}

	// Load the created transaction with relations
	if err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("User").
		Preload("Items").
		Preload("Items.Service").
		First(&transaction, transaction.ID).Error; err != nil {
		return entities.Transaction{}, errors.Wrap(err, "failed to load created transaction")
	}

	return transaction, nil
}

func (r *transactionPostgresRepository) GetList(ctx context.Context) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	query := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("User").
		Preload("Items").
		Preload("Items.Service")

	if err := query.Order("created_at DESC").Find(&transactions).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get transactions")
	}
	return transactions, nil
}

func (r *transactionPostgresRepository) FindById(ctx context.Context, transactionId uint) (entities.Transaction, error) {
	var transaction entities.Transaction

	if err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("User").
		Preload("Items").
		Preload("Items.Service").
		Where("id = ?", transactionId).
		First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Transaction{}, fmt.Errorf("transaction with id %d not found", transactionId)
		}
		return entities.Transaction{}, errors.Wrap(err, "failed to find transaction")
	}

	return transaction, nil
}

func (r *transactionPostgresRepository) Update(ctx context.Context, transactionId uint, data entities.Transaction) (entities.Transaction, error) {
	var existing entities.Transaction

	// Check if transaction exists
	if err := r.db.WithContext(ctx).First(&existing, transactionId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Transaction{}, fmt.Errorf("transaction with id %d not found", transactionId)
		}
		return entities.Transaction{}, errors.Wrap(err, "failed to find transaction")
	}

	// Update only provided fields
	updateData := entities.Transaction{}
	if data.OrderStatus != "" {
		updateData.OrderStatus = data.OrderStatus
	}
	if data.PaymentStatus != "" {
		updateData.PaymentStatus = data.PaymentStatus
	}
	if data.AmountPaid != 0 {
		updateData.AmountPaid = data.AmountPaid
	}

	if err := r.db.WithContext(ctx).Model(&existing).Updates(updateData).Error; err != nil {
		return entities.Transaction{}, errors.Wrap(err, "failed to update transaction")
	}

	// Load updated transaction with relations
	if err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("User").
		Preload("Items").
		Preload("Items.Service").
		First(&existing, transactionId).Error; err != nil {
		return entities.Transaction{}, errors.Wrap(err, "failed to load updated transaction")
	}

	return existing, nil
}
