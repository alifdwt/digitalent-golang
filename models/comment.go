package models

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PhotoID   uint      `json:"photo_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"size:200;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	User      User
	Photo     Photo
}
