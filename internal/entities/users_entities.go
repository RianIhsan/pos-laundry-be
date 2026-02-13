package entities

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null;type:varchar(50)" json:"username"`
	Password  string    `gorm:"not null;type:text" json:"-"` // "-" agar password tidak pernah bocor ke JSON
	Name      string    `gorm:"not null;type:varchar(100)" json:"name"`
	Role      string    `gorm:"type:varchar(20);default:'owner'" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (User) TableName() string {
	return "master_users"
}
