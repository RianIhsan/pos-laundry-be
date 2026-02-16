package repository

import (
	"context"
	"fmt"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/customers"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type customerPostgresRepository struct {
	db *gorm.DB
}

func NewCustomerPostgresRepository(db *gorm.DB) customers.CustomerRepositoryInterface {
	return &customerPostgresRepository{db: db}
}

func (c *customerPostgresRepository) Create(ctx context.Context, entity entities.Customer) (entities.Customer, error) {
	if err := c.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return entities.Customer{}, err
	}
	return entity, nil
}

func (c *customerPostgresRepository) GetList(ctx context.Context) ([]entities.Customer, error) {
	var customers []entities.Customer
	query := c.db.WithContext(ctx).Model(&entities.Customer{}).
		Preload("Transactions")

	if err := query.Order("created_at ASC").Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerPostgresRepository) FindById(ctx context.Context, customerId uint) (entities.Customer, error) {
	var customer entities.Customer
	DB := c.db.WithContext(ctx)

	if err := DB.Where("id = ?", customerId).Take(&customer).Error; err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}

func (c *customerPostgresRepository) Update(ctx context.Context, id uint, data entities.Customer) error {
	var existing entities.Customer
	if err := c.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("customer with id %d not found", id)
		}
		return fmt.Errorf("cannot update customer with id %d: %w", id, err)
	}
	if err := c.db.WithContext(ctx).Model(&existing).Updates(data).Error; err != nil {
		return fmt.Errorf("cannot update customer with id %d: %w", id, err)
	}
	return nil
}

func (c *customerPostgresRepository) DeleteCustomer(ctx context.Context, customerId uint) error {
	// Start transaction to ensure atomicity
	tx := c.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Delete all transactions related to this customer
	// This will also cascade delete transaction items and logs due to OnDelete:CASCADE constraint
	if err := tx.Where("customer_id = ?", customerId).Delete(&entities.Transaction{}).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to delete customer transactions")
	}

	// Delete the customer
	if err := tx.Delete(&entities.Customer{}, customerId).Error; err != nil {
		tx.Rollback()
		return errors.Wrap(err, "failed to delete customer")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.Wrap(err, "failed to commit delete customer transaction")
	}

	return nil
}
