package repository

import "github.com/alifdwt/digitalent-golang/models"

type OrderRepository interface {
	CreateOrder(request *models.Order) (*models.Order, error)
	GetAllOrder() (*[]models.Order, error)
	UpdateOrder(orderId uint, request *models.Order) (*models.Order, error)
	DeleteOrder(orderId uint) error
}
