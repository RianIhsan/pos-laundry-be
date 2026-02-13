package entities

import "time"

type TransactionLog struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID uint      `json:"transaction_id"`
	Activity      string    `gorm:"not null;type:text" json:"activity"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
