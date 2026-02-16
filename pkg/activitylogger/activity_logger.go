package activitylogger

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"gorm.io/gorm"
)

type ActivityLogger struct {
	db *gorm.DB
}

func NewActivityLogger(db *gorm.DB) *ActivityLogger {
	return &ActivityLogger{db: db}
}

func (al *ActivityLogger) LogActivity(ctx context.Context, userID uint, action string, targetType string, targetID string, description string) error {
	activityLog := entities.ActivityLog{
		UserID:      userID,
		Action:      action,
		TargetType:  targetType,
		TargetID:    targetID,
		Description: description,
	}

	if err := al.db.WithContext(ctx).Create(&activityLog).Error; err != nil {
		return err
	}
	return nil
}

// Helper methods untuk berbagai aksi
func (al *ActivityLogger) LogTransactionCreated(ctx context.Context, userID uint, invoiceNo string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Membuat pesanan",
		"TRANSACTION",
		invoiceNo,
		"Pesanan baru dibuat",
	)
}

func (al *ActivityLogger) LogTransactionStatusUpdated(ctx context.Context, userID uint, invoiceNo string, newOrderStatus string, newPaymentStatus string) error {
	description := "Status diperbarui"
	if newOrderStatus != "" {
		description += " - Order: " + newOrderStatus
	}
	if newPaymentStatus != "" {
		description += " - Payment: " + newPaymentStatus
	}

	return al.LogActivity(
		ctx,
		userID,
		"Mengubah status pesanan",
		"TRANSACTION",
		invoiceNo,
		description,
	)
}

func (al *ActivityLogger) LogCustomerCreated(ctx context.Context, userID uint, customerID string, customerName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Menambah customer baru",
		"CUSTOMER",
		customerID,
		"Customer: "+customerName,
	)
}

func (al *ActivityLogger) LogCustomerUpdated(ctx context.Context, userID uint, customerID string, customerName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Mengubah data customer",
		"CUSTOMER",
		customerID,
		"Customer: "+customerName,
	)
}

func (al *ActivityLogger) LogCustomerDeleted(ctx context.Context, userID uint, customerID string, customerName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Menghapus customer",
		"CUSTOMER",
		customerID,
		"Customer: "+customerName,
	)
}

func (al *ActivityLogger) LogServiceCreated(ctx context.Context, userID uint, serviceID string, serviceName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Menambah layanan baru",
		"SERVICE",
		serviceID,
		"Service: "+serviceName,
	)
}

func (al *ActivityLogger) LogServiceUpdated(ctx context.Context, userID uint, serviceID string, serviceName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Mengubah layanan",
		"SERVICE",
		serviceID,
		"Service: "+serviceName,
	)
}

func (al *ActivityLogger) LogServiceDeleted(ctx context.Context, userID uint, serviceID string, serviceName string) error {
	return al.LogActivity(
		ctx,
		userID,
		"Menghapus layanan",
		"SERVICE",
		serviceID,
		"Service: "+serviceName,
	)
}
