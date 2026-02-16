package dto

type CreateInventoryRequest struct {
	Name          string  `json:"name" binding:"required"`
	Description   string  `json:"description"`
	Category      string  `json:"category" binding:"required"` // DETERGENT, FRAGRANCE, PACKAGING, etc
	CurrentStock  float64 `json:"current_stock" binding:"required,min=0"`
	MaxStock      float64 `json:"max_stock" binding:"required,min=0"`
	CriticalLevel float64 `json:"critical_level" binding:"required,min=0"`
	Unit          string  `json:"unit" binding:"required"` // Liter, Kg, Pcs, etc
}

type UpdateInventoryRequest struct {
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Category      *string  `json:"category"`
	CurrentStock  *float64 `json:"current_stock" binding:"omitempty,min=0"`
	MaxStock      *float64 `json:"max_stock" binding:"omitempty,min=0"`
	CriticalLevel *float64 `json:"critical_level" binding:"omitempty,min=0"`
	Unit          *string  `json:"unit"`
}
