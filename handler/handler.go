package handler

import (
	_ "github.com/alifdwt/digitalent-golang/docs"
	"github.com/alifdwt/digitalent-golang/domain/responses"
	"github.com/alifdwt/digitalent-golang/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services service.Service
	router   *gin.Engine
}

func NewHandler(services service.Service) (*Handler, error) {
	handler := &Handler{
		services: services,
	}

	handler.setupRouter()
	return handler, nil
}

func (h *Handler) setupRouter() {
	h.router = gin.Default()

	h.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.router.POST("/orders", h.createOrder)
	h.router.GET("/orders", h.getAllOrder)
	h.router.PUT("/orders/:orderId", h.updateOrder)
	h.router.DELETE("/orders/:orderId", h.deleteOrder)
}

func (h *Handler) Start(address string) error {
	return h.router.Run(address)
}

func errorResponse(err error) responses.ErrorMessage {
	return responses.ErrorMessage{
		Message: err.Error(),
	}
}
