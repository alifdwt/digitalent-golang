package mapper

import (
	"github.com/alifdwt/digitalent-golang/domain/responses/order"
	"github.com/alifdwt/digitalent-golang/models"
)

type OrderMapping interface {
	ToOrderResponse(request *models.Order) *order.OrderResponse
	ToOrderResponses(requests *[]models.Order) []order.OrderResponse
}
