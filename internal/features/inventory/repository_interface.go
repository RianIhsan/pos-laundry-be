package inventory

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
)

type InventoryRepositoryInterface interface {
	GetAllItems(ctx context.Context) ([]entities.InventoryItem, error)
	GetDetailByID(ctx context.Context, id uint) (entities.InventoryItem, error)
	Create(ctx context.Context, item entities.InventoryItem) (entities.InventoryItem, error)
	Update(ctx context.Context, id uint, item entities.InventoryItem) (entities.InventoryItem, error)
	Delete(ctx context.Context, id uint) error
}
