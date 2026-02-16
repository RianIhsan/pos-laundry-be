package repository

import (
	"context"
	"time"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/dashboard"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dashboardPostgresRepository struct {
	db *gorm.DB
}

func NewDashboardPostgresRepository(db *gorm.DB) dashboard.DashboardRepositoryInterface {
	return &dashboardPostgresRepository{db: db}
}

func (r *dashboardPostgresRepository) GetTotalOmzet(ctx context.Context) (float64, error) {
	var total float64
	today := time.Now().Format("2006-01-02")

	if err := r.db.WithContext(ctx).
		Model(&entities.Transaction{}).
		Where("DATE(created_at) = ?", today).
		Select("COALESCE(SUM(total_price), 0) as total").
		Row().
		Scan(&total); err != nil {
		return 0, errors.Wrap(err, "failed to get total omzet")
	}
	return total, nil
}

func (r *dashboardPostgresRepository) GetNewMembersCount(ctx context.Context) (int64, error) {
	var count int64
	today := time.Now().Format("2006-01-02")

	if err := r.db.WithContext(ctx).
		Model(&entities.Customer{}).
		Where("DATE(created_at) = ?", today).
		Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "failed to get new members count")
	}
	return count, nil
}

func (r *dashboardPostgresRepository) GetActiveProductionCount(ctx context.Context) (int64, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&entities.Transaction{}).
		Where("order_status IN ?", []string{"WASHING", "DRYING"}).
		Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "failed to get active production count")
	}
	return count, nil
}

func (r *dashboardPostgresRepository) GetReadyToPickupCount(ctx context.Context) (int64, error) {
	var count int64

	if err := r.db.WithContext(ctx).
		Model(&entities.Transaction{}).
		Where("order_status = ?", "READY_TO_PICKUP").
		Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "failed to get ready to pickup count")
	}
	return count, nil
}

func (r *dashboardPostgresRepository) GetUnpaidRevenue(ctx context.Context) (float64, error) {
	var total float64

	if err := r.db.WithContext(ctx).
		Model(&entities.Transaction{}).
		Where("payment_status = ?", "UNPAID").
		Select("COALESCE(SUM(total_price), 0) as total").
		Row().
		Scan(&total); err != nil {
		return 0, errors.Wrap(err, "failed to get unpaid revenue")
	}
	return total, nil
}

func (r *dashboardPostgresRepository) GetProductionQueues(ctx context.Context, limit int) ([]entities.Transaction, error) {
	var transactions []entities.Transaction

	if err := r.db.WithContext(ctx).
		Preload("Customer").
		Preload("User").
		Preload("Items").
		Preload("Items.Service").
		Where("order_status IN ?", []string{"WASHING", "DRYING"}).
		Order("created_at ASC").
		Limit(limit).
		Find(&transactions).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get production queues")
	}
	return transactions, nil
}

func (r *dashboardPostgresRepository) GetPopularServices(ctx context.Context, limit int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	if err := r.db.WithContext(ctx).
		Table("transaction_items ti").
		Select("s.name as name, COUNT(ti.id) as total_orders, SUM(ti.subtotal) as total_revenue").
		Joins("JOIN master_services s ON ti.service_id = s.id").
		Group("ti.service_id, s.name").
		Order("total_orders DESC").
		Limit(limit).
		Scan(&results).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get popular services")
	}
	return results, nil
}

func (r *dashboardPostgresRepository) GetActivityLogs(ctx context.Context, limit int) ([]entities.ActivityLog, error) {
	var logs []entities.ActivityLog

	if err := r.db.WithContext(ctx).
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Find(&logs).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get activity logs")
	}
	return logs, nil
}

func (r *dashboardPostgresRepository) GetActivityLogsPaginated(ctx context.Context, limit int, offset int) ([]entities.ActivityLog, error) {
	var logs []entities.ActivityLog

	if err := r.db.WithContext(ctx).
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&logs).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get activity logs paginated")
	}
	return logs, nil
}

func (r *dashboardPostgresRepository) CountActivityLogs(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&entities.ActivityLog{}).
		Count(&count).Error; err != nil {
		return 0, errors.Wrap(err, "failed to count activity logs")
	}
	return count, nil
}
