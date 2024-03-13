package service

import (
	"github.com/alifdwt/digitalent-golang/domain/responses/order"
	"github.com/alifdwt/digitalent-golang/models"
)

type OrderService interface {
	CreateOrder(request *models.Order) (*order.OrderResponse, error)
	GetAllOrder() (*[]order.OrderResponse, error)
	UpdateOrder(orderId uint, request *models.Order) (*order.OrderResponse, error)
	DeleteOrder(orderId uint) error
}
