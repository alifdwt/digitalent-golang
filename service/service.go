package service

import (
	"github.com/alifdwt/digitalent-golang/mapper"
	"github.com/alifdwt/digitalent-golang/repository"
)

type Service struct {
	Order OrderService
}

type Deps struct {
	Repositories *repository.Repositories
	Mapper       mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Order: NewOrderService(deps.Repositories.Order, deps.Mapper.OrderMapper),
	}
}
