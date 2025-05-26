package models

import (
	"time"
)

type Post struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"` 
	UpdatedAt   time.Time `json:"updated_at"`
}
