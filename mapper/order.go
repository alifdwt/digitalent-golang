package mapper

import (
	"github.com/alifdwt/digitalent-golang/domain/responses/order"
	"github.com/alifdwt/digitalent-golang/models"
)

type orderMapper struct {
}

func NewOrderMapper() *orderMapper {
	return &orderMapper{}
}

func (m *orderMapper) ToOrderResponse(request *models.Order) *order.OrderResponse {
	return &order.OrderResponse{
		ID:           request.ID,
		CustomerName: request.CustomerName,
		OrderedAt:    request.OrderedAt,
		Items:        m.ToItemsResponses(request.Items),
	}
}

func (m *orderMapper) ToOrderResponses(requests *[]models.Order) []order.OrderResponse {
	var responses []order.OrderResponse

	for _, request := range *requests {
		response := order.OrderResponse{
			ID:           request.ID,
			CustomerName: request.CustomerName,
			OrderedAt:    request.OrderedAt,
			Items:        m.ToItemsResponses(request.Items),
		}
		responses = append(responses, response)
	}

	return responses
}

func (m *orderMapper) ToItemsResponses(requests []models.Items) []order.ItemResponse {
	var responses []order.ItemResponse

	for _, request := range requests {
		response := order.ItemResponse{
			ItemCode:    request.Code,
			Description: request.Description,
			Quantity:    request.Quantity,
		}

		responses = append(responses, response)
	}

	return responses
}
