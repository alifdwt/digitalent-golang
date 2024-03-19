package handler

import (
	_ "mygram-api/docs"
	"mygram-api/domain/responses"
	"mygram-api/pkg/token"
	"mygram-api/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services   *service.Service
	tokenMaker token.Maker
}

func NewHandler(services *service.Service, tokenMaker token.Maker) *Handler {
	return &Handler{
		services:   services,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *gin.Engine) {
	h.initAuthGroup(router)
	h.initUserGroup(router)
}

func errorResponse(err error) responses.ErrorMessage {
	return responses.ErrorMessage{
		Message: err.Error(),
	}
}
