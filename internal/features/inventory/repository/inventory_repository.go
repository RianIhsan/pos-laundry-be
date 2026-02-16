package repository

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type inventoryPostgresRepository struct {
	db *gorm.DB
}

func NewInventoryPostgresRepository(db *gorm.DB) inventory.InventoryRepositoryInterface {
	return &inventoryPostgresRepository{db: db}
}

func (r *inventoryPostgresRepository) GetAllItems(ctx context.Context) ([]entities.InventoryItem, error) {
	var items []entities.InventoryItem

	if err := r.db.WithContext(ctx).
		Order("current_stock ASC").
		Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get inventory items")
	}
	return items, nil
}

func (r *inventoryPostgresRepository) GetDetailByID(ctx context.Context, id uint) (entities.InventoryItem, error) {
	var item entities.InventoryItem

	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.InventoryItem{}, errors.New("inventory item not found")
		}
		return entities.InventoryItem{}, errors.Wrap(err, "failed to get inventory item")
	}
	return item, nil
}

func (r *inventoryPostgresRepository) Create(ctx context.Context, item entities.InventoryItem) (entities.InventoryItem, error) {
	if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
		return entities.InventoryItem{}, errors.Wrap(err, "failed to create inventory item")
	}
	return item, nil
}

func (r *inventoryPostgresRepository) Update(ctx context.Context, id uint, item entities.InventoryItem) (entities.InventoryItem, error) {
	if err := r.db.WithContext(ctx).Model(&entities.InventoryItem{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return entities.InventoryItem{}, errors.Wrap(err, "failed to update inventory item")
	}

	var updated entities.InventoryItem
	if err := r.db.WithContext(ctx).First(&updated, id).Error; err != nil {
		return entities.InventoryItem{}, errors.Wrap(err, "failed to fetch updated inventory item")
	}
	return updated, nil
}

func (r *inventoryPostgresRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entities.InventoryItem{}, id).Error; err != nil {
		return errors.Wrap(err, "failed to delete inventory item")
	}
	return nil
}
