package dto

import "time"

type InventoryItemDTO struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Category        string    `json:"category"`
	CurrentStock    float64   `json:"current_stock"`
	MaxStock        float64   `json:"max_stock"`
	CriticalLevel   float64   `json:"critical_level"`
	Unit            string    `json:"unit"`
	StockPercentage int32     `json:"stock_percentage"`
	Status          string    `json:"status"`
	LastRestockDate time.Time `json:"last_restock_date"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateInventoryResponse struct {
	ID uint `json:"id"`
}

type InventoryAlertDTO struct {
	ItemName        string `json:"item_name"`
	StockPercentage int32  `json:"stock_percentage"`
	Status          string `json:"status"` // CRITICAL, WARNING, SAFE
}
