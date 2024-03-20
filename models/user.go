package models

import "time"

type User struct {
	ID              uint          `json:"id" gorm:"primaryKey"`
	Username        string        `json:"username" gorm:"size:50;unique;not null"`
	Email           string        `json:"email" gorm:"size:150;unique;not null"`
	Password        string        `json:"password" gorm:"type:text;not null"`
	Age             uint          `json:"age" gorm:"not null"`
	ProfileImageURL string        `json:"profile_image_url"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Comments        []Comment     `gorm:"constraint:OnDelete:CASCADE"`
	Photos          []Photo       `gorm:"constraint:OnDelete:CASCADE"`
	SocialMedia     []SocialMedia `gorm:"constraint:OnDelete:CASCADE"`
}
