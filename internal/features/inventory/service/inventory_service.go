package service

import (
	"context"
	"fmt"
	"math"

	"github.com/RianIhsan/pos-laundry-be/internal/entities"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory"
	"github.com/RianIhsan/pos-laundry-be/internal/features/inventory/dto"
	"github.com/pkg/errors"
)

type inventoryService struct {
	cfg           *ServiceConfig
	inventoryRepo inventory.InventoryRepositoryInterface
}

func NewInventoryService(cfg *ServiceConfig) inventory.InventoryServiceInterface {
	return &inventoryService{
		cfg:           cfg,
		inventoryRepo: cfg.InventoryRepoInterface,
	}
}

func (s *inventoryService) GetAlerts(ctx context.Context) ([]dto.InventoryAlertDTO, error) {
	items, err := s.inventoryRepo.GetAllItems(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get inventory items")
	}

	alerts := make([]dto.InventoryAlertDTO, 0)
	for _, item := range items {
		// Calculate stock percentage
		stockPercentage := int32(0)
		if item.MaxStock > 0 {
			stockPercentage = int32(math.Round((item.CurrentStock / item.MaxStock) * 100))
		}

		// Determine status
		status := "SAFE"
		if item.CurrentStock <= item.CriticalLevel {
			status = "CRITICAL"
		} else if item.CurrentStock <= (item.CriticalLevel * 2) {
			status = "WARNING"
		}

		alert := dto.InventoryAlertDTO{
			ItemName:        item.Name,
			StockPercentage: stockPercentage,
			Status:          status,
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}

func (s *inventoryService) Create(ctx context.Context, userID uint, req dto.CreateInventoryRequest) (dto.CreateInventoryResponse, error) {
	item := entities.InventoryItem{
		Name:          req.Name,
		Description:   req.Description,
		Category:      req.Category,
		CurrentStock:  req.CurrentStock,
		MaxStock:      req.MaxStock,
		CriticalLevel: req.CriticalLevel,
		Unit:          req.Unit,
	}

	created, err := s.inventoryRepo.Create(ctx, item)
	if err != nil {
		return dto.CreateInventoryResponse{}, errors.Wrap(err, "failed to create inventory item")
	}

	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogActivity(ctx, userID, "Membuat inventory item", "INVENTORY", fmt.Sprintf("%d", created.ID), "")
	}

	return dto.CreateInventoryResponse{ID: created.ID}, nil
}

func (s *inventoryService) List(ctx context.Context) ([]dto.InventoryItemDTO, error) {
	items, err := s.inventoryRepo.GetAllItems(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get inventory items")
	}

	res := make([]dto.InventoryItemDTO, 0)
	for _, it := range items {
		percentage := int32(0)
		if it.MaxStock > 0 {
			percentage = int32(math.Round((it.CurrentStock / it.MaxStock) * 100))
		}
		status := "SAFE"
		if it.CurrentStock <= it.CriticalLevel {
			status = "CRITICAL"
		} else if it.CurrentStock <= (it.CriticalLevel * 2) {
			status = "WARNING"
		}
		res = append(res, dto.InventoryItemDTO{
			ID:              it.ID,
			Name:            it.Name,
			Description:     it.Description,
			Category:        it.Category,
			CurrentStock:    it.CurrentStock,
			MaxStock:        it.MaxStock,
			CriticalLevel:   it.CriticalLevel,
			Unit:            it.Unit,
			StockPercentage: percentage,
			Status:          status,
			LastRestockDate: it.LastRestockDate,
			CreatedAt:       it.CreatedAt,
		})
	}
	return res, nil
}

func (s *inventoryService) GetByID(ctx context.Context, id uint) (dto.InventoryItemDTO, error) {
	it, err := s.inventoryRepo.GetDetailByID(ctx, id)
	if err != nil {
		return dto.InventoryItemDTO{}, errors.Wrap(err, "failed to get inventory item")
	}
	percentage := int32(0)
	if it.MaxStock > 0 {
		percentage = int32(math.Round((it.CurrentStock / it.MaxStock) * 100))
	}
	status := "SAFE"
	if it.CurrentStock <= it.CriticalLevel {
		status = "CRITICAL"
	} else if it.CurrentStock <= (it.CriticalLevel * 2) {
		status = "WARNING"
	}
	return dto.InventoryItemDTO{
		ID:              it.ID,
		Name:            it.Name,
		Description:     it.Description,
		Category:        it.Category,
		CurrentStock:    it.CurrentStock,
		MaxStock:        it.MaxStock,
		CriticalLevel:   it.CriticalLevel,
		Unit:            it.Unit,
		StockPercentage: percentage,
		Status:          status,
		LastRestockDate: it.LastRestockDate,
		CreatedAt:       it.CreatedAt,
	}, nil
}

func (s *inventoryService) Update(ctx context.Context, userID uint, id uint, req dto.UpdateInventoryRequest) (dto.InventoryItemDTO, error) {
	existing, err := s.inventoryRepo.GetDetailByID(ctx, id)
	if err != nil {
		return dto.InventoryItemDTO{}, errors.Wrap(err, "failed to fetch existing inventory item")
	}
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.Category != nil {
		existing.Category = *req.Category
	}
	if req.CurrentStock != nil {
		existing.CurrentStock = *req.CurrentStock
	}
	if req.MaxStock != nil {
		existing.MaxStock = *req.MaxStock
	}
	if req.CriticalLevel != nil {
		existing.CriticalLevel = *req.CriticalLevel
	}
	if req.Unit != nil {
		existing.Unit = *req.Unit
	}

	updated, err := s.inventoryRepo.Update(ctx, id, existing)
	if err != nil {
		return dto.InventoryItemDTO{}, errors.Wrap(err, "failed to update inventory item")
	}

	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogActivity(ctx, userID, "Mengubah inventory item", "INVENTORY", fmt.Sprintf("%d", updated.ID), "")
	}

	percentage := int32(0)
	if updated.MaxStock > 0 {
		percentage = int32(math.Round((updated.CurrentStock / updated.MaxStock) * 100))
	}
	status := "SAFE"
	if updated.CurrentStock <= updated.CriticalLevel {
		status = "CRITICAL"
	} else if updated.CurrentStock <= (updated.CriticalLevel * 2) {
		status = "WARNING"
	}
	return dto.InventoryItemDTO{
		ID:              updated.ID,
		Name:            updated.Name,
		Description:     updated.Description,
		Category:        updated.Category,
		CurrentStock:    updated.CurrentStock,
		MaxStock:        updated.MaxStock,
		CriticalLevel:   updated.CriticalLevel,
		Unit:            updated.Unit,
		StockPercentage: percentage,
		Status:          status,
		LastRestockDate: updated.LastRestockDate,
		CreatedAt:       updated.CreatedAt,
	}, nil
}

func (s *inventoryService) Delete(ctx context.Context, userID uint, id uint) error {
	if err := s.inventoryRepo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete inventory item")
	}
	if s.cfg.ActivityLogger != nil {
		_ = s.cfg.ActivityLogger.LogActivity(ctx, userID, "Menghapus inventory item", "INVENTORY", fmt.Sprintf("%d", id), "")
	}
	return nil
}
