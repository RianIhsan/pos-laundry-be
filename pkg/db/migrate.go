package db

import (
	"fmt"
	"github.com/RianIhsan/pos-laundry-be/internal/entities"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Customer{},
		&entities.Service{},
		&entities.Transaction{},
		&entities.TransactionItem{},
		&entities.TransactionLog{},
	)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}
