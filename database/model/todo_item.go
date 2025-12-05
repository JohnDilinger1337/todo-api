package model

import "time"

// import 	"golang.org/x/crypto/bcrypt"

type TodoItem struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`
	Description string       `gorm:"not null" json:"description"`
	Status      string       `gorm:"not null" json:"status"`
	UserID      uint         `gorm:"not null" json:"user_id"`
	CategoryID  uint         `gorm:"not null" json:"category_id"`
	Category    TodoCategory `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
