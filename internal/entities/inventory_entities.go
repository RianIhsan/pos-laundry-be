package entities

import "time"

type InventoryItem struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null;type:varchar(100)" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Category        string    `gorm:"not null;type:varchar(50)" json:"category"` // DETERGENT, FRAGRANCE, PACKAGING, etc
	CurrentStock    float64   `gorm:"not null;type:decimal(10,2)" json:"current_stock"`
	MaxStock        float64   `gorm:"not null;type:decimal(10,2)" json:"max_stock"`
	CriticalLevel   float64   `gorm:"not null;type:decimal(10,2)" json:"critical_level"`
	Unit            string    `gorm:"not null;type:varchar(20)" json:"unit"` // Liter, Kg, Pcs, etc
	LastRestockDate time.Time `gorm:"type:timestamp" json:"last_restock_date"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (InventoryItem) TableName() string {
	return "master_inventory_items"
}
