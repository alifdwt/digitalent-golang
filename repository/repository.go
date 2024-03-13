package repository

import "gorm.io/gorm"

type Repositories struct {
	Order OrderRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Order: NewOrderRepository(db),
	}
}
