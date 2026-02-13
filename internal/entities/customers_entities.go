package entities

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;type:varchar(100)" json:"name"`
	Phone     string    `gorm:"unique;not null;type:varchar(20)" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Points    int       `gorm:"default:0" json:"points"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// Relasi: Satu customer bisa punya banyak transaksi
	Transactions []Transaction `gorm:"foreignKey:CustomerID" json:"transactions,omitempty"`
}

func (Customer) TableName() string {
	return "master_customers"
}
