package entities

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ServiceID     uint    `json:"service_id"`
	Service       Service `gorm:"foreignKey:ServiceID" json:"service"`
	Qty           float64 `gorm:"not null;type:decimal(10,2)" json:"qty"`
	PriceAtTime   float64 `gorm:"not null;type:decimal(12,2)" json:"price_at_time"`
	Notes         string  `gorm:"type:text" json:"note"`
	Subtotal      float64 `gorm:"not null;type:decimal(12,2)" json:"subtotal"`
}
