package models

type Items struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Code        string `json:"itemCode" gorm:"size:10"`
	Description string `json:"description" gorm:"size:50"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
