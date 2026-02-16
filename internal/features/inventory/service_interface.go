package inventory

import (
	"context"

	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory/dto"
)

type InventoryServiceInterface interface {
	GetAlerts(ctx context.Context) ([]dto.InventoryAlertDTO, error)
	Create(ctx context.Context, userID uint, req dto.CreateInventoryRequest) (dto.CreateInventoryResponse, error)
	List(ctx context.Context) ([]dto.InventoryItemDTO, error)
	GetByID(ctx context.Context, id uint) (dto.InventoryItemDTO, error)
	Update(ctx context.Context, userID uint, id uint, req dto.UpdateInventoryRequest) (dto.InventoryItemDTO, error)
	Delete(ctx context.Context, userID uint, id uint) error
}
