package service

import (
	"context"
	"fmt"
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions"
	"github.com/RianIhsan/pos-laundry-be/internal/features/transactions/dto"
	"github.com/pkg/errors"
)

type transactionService struct {
	cfg             *ServiceConfig
	transactionRepo transactions.TransactionRepositoryInterface
}

func NewTransactionService(cfg *ServiceConfig) transactions.TransactionServiceInterface {
	return &transactionService{
		cfg:             cfg,
		transactionRepo: cfg.TransactionRepoInterface,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest, userID uint) (dto.TransactionResponse, error) {
	// Prepare request data
	req.PrepareCreate()

	// Validate items
	if len(req.Items) == 0 {
		return dto.TransactionResponse{}, errors.New("transaction must have at least one item")
	}

	// Generate invoice number
	invoiceNo := s.generateInvoiceNumber()

	// Convert to transaction entity
	transactionEntity := dto.ConvertToTransactionEntity(req, invoiceNo, userID)

	// Convert items to entities
	var itemEntities []entities.TransactionItem
	for _, item := range req.Items {
		itemEntity := dto.ConvertToTransactionItemEntity(0, item) // TransactionID will be set during create
		itemEntities = append(itemEntities, itemEntity)
	}

	// Create transaction and items
	created, err := s.transactionRepo.Create(ctx, transactionEntity, itemEntities)
	if err != nil {
		return dto.TransactionResponse{}, errors.Wrap(err, "failed to create transaction")
	}

	// Log activity
	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogTransactionCreated(ctx, userID, created.InvoiceNo)
	}

	return dto.ToTransactionResponse(created), nil
}

func (s *transactionService) GetList(ctx context.Context) ([]dto.TransactionResponse, error) {
	transactions, err := s.transactionRepo.GetList(ctx)
	if err != nil {
		return []dto.TransactionResponse{}, errors.Wrap(err, "failed to get transactions")
	}

	response := make([]dto.TransactionResponse, 0)
	for _, tx := range transactions {
		response = append(response, dto.ToTransactionResponse(tx))
	}

	return response, nil
}

func (s *transactionService) GetById(ctx context.Context, transactionId uint) (dto.TransactionResponse, error) {
	transaction, err := s.transactionRepo.FindById(ctx, transactionId)
	if err != nil {
		return dto.TransactionResponse{}, errors.Wrap(err, "failed to find transaction")
	}

	return dto.ToTransactionResponse(transaction), nil
}

func (s *transactionService) UpdateStatus(ctx context.Context, transactionId uint, userID uint, req dto.UpdateTransactionStatusRequest) (dto.TransactionResponse, error) {
	// Prepare request data
	req.PrepareUpdate()

	// Validate that at least one status is provided
	if req.OrderStatus == "" && req.PaymentStatus == "" {
		return dto.TransactionResponse{}, errors.New("at least one status must be provided")
	}

	// Get transaction before update to log
	transactionBefore, err := s.transactionRepo.FindById(ctx, transactionId)
	if err != nil {
		return dto.TransactionResponse{}, errors.Wrap(err, "failed to find transaction")
	}

	// Create update entity
	updateData := entities.Transaction{
		OrderStatus:   req.OrderStatus,
		PaymentStatus: req.PaymentStatus,
	}

	// Update transaction
	updated, err := s.transactionRepo.Update(ctx, transactionId, updateData)
	if err != nil {
		return dto.TransactionResponse{}, errors.Wrap(err, "failed to update transaction status")
	}

	// Log activity
	if s.cfg.ActivityLogger != nil {
		// Use current status if not changed
		newOrderStatus := req.OrderStatus
		if newOrderStatus == "" {
			newOrderStatus = transactionBefore.OrderStatus
		}
		newPaymentStatus := req.PaymentStatus
		if newPaymentStatus == "" {
			newPaymentStatus = transactionBefore.PaymentStatus
		}
		_ = s.cfg.ActivityLogger.LogTransactionStatusUpdated(
			ctx,
			userID,
			transactionBefore.InvoiceNo,
			newOrderStatus,
			newPaymentStatus,
		)
	}

	return dto.ToTransactionResponse(updated), nil
}

// generateInvoiceNumber generates a unique invoice number
// Format: INV-YYMMDD-XXXXX (where XXXXX is timestamp based counter)
func (s *transactionService) generateInvoiceNumber() string {
	now := time.Now()
	timestamp := now.UnixNano()

	return fmt.Sprintf("INV-%s-%d",
		now.Format("060102"), // YYMMDD
		timestamp%100000,     // Last 5 digits of timestamp for uniqueness
	)
}
