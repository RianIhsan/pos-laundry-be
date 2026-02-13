package entities

import "time"

type Service struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;type:varchar(100)" json:"name"`
	Category  string    `gorm:"not null;type:varchar(50)" json:"category"`
	Price     float64   `gorm:"not null;type:decimal(12,2)" json:"price"`
	Unit      string    `gorm:"not null;type:varchar(10)" json:"unit"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Service) TableName() string {
	return "master_services"
}
