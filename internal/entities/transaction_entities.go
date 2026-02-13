package entities

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	InvoiceNo     string    `gorm:"unique;not null;type:varchar(20)" json:"invoice_no"`
	CustomerID    uint      `json:"customer_id"`
	Customer      Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	UserID        uint      `json:"user_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	TotalPrice    float64   `gorm:"not null;type:decimal(12,2)" json:"total_price"`
	PaymentMethod string    `gorm:"not null;type:varchar(20)" json:"payment_method"` // CASH, QRIS
	PaymentStatus string    `gorm:"type:varchar(20);default:'UNPAID'" json:"payment_status"`
	OrderStatus   string    `gorm:"type:varchar(20);default:'WASHING'" json:"order_status"`
	AmountPaid    float64   `gorm:"type:decimal(12,2)" json:"amount_paid"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// Relasi Has Many
	Items []TransactionItem `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"items"`
	Logs  []TransactionLog  `gorm:"foreignKey:TransactionID;constraint:OnDelete:CASCADE" json:"logs"`
}

func (Transaction) TableName() string {
	return "master_transactions"
}
