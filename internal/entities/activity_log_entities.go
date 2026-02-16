package entities

import "time"

type ActivityLog struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	Action      string    `gorm:"not null;type:varchar(100)" json:"action"`
	TargetType  string    `gorm:"not null;type:varchar(50)" json:"target_type"` // TRANSACTION, CUSTOMER, SERVICE, etc
	TargetID    string    `gorm:"not null;type:varchar(50)" json:"target_id"`   // Invoice number, customer ID, etc
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (ActivityLog) TableName() string {
	return "activity_logs"
}
