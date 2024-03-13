package service

import (
	"fmt"

	"github.com/alifdwt/digitalent-golang/domain/responses/order"
	"github.com/alifdwt/digitalent-golang/mapper"
	"github.com/alifdwt/digitalent-golang/models"
	"github.com/alifdwt/digitalent-golang/repository"
)

type orderService struct {
	repository repository.OrderRepository
	mapper     mapper.OrderMapping
}

func NewOrderService(orderRepository repository.OrderRepository, mapper mapper.OrderMapping) *orderService {
	return &orderService{
		repository: orderRepository,
		mapper:     mapper,
	}
}

func (s *orderService) CreateOrder(request *models.Order) (*order.OrderResponse, error) {
	res, err := s.repository.CreateOrder(request)
	if err != nil {
		fmt.Printf("error while creating order: %s", err.Error())
		return nil, err
	}

	mapper := s.mapper.ToOrderResponse(res)

	return mapper, nil
}

func (s *orderService) GetAllOrder() (*[]order.OrderResponse, error) {
	res, err := s.repository.GetAllOrder()
	if err != nil {
		fmt.Printf("error while retrieve all orders: %s", err.Error())
		return nil, err
	}

	mapper := s.mapper.ToOrderResponses(res)

	return &mapper, nil
}

func (s *orderService) UpdateOrder(orderId uint, request *models.Order) (*order.OrderResponse, error) {
	res, err := s.repository.UpdateOrder(orderId, request)
	if err != nil {
		fmt.Printf("error while updating order: %s", err.Error())
		return nil, err
	}

	mapper := s.mapper.ToOrderResponse(res)

	return mapper, nil
}

func (s *orderService) DeleteOrder(orderId uint) error {
	return s.repository.DeleteOrder(orderId)
}
