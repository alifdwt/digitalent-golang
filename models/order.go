package models

import (
	"time"
)

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customerName" gorm:"size:50"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Items   `json:"items" gorm:"foreignKey:OrderID"`
}
