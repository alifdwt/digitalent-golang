package handler

import (
	"net/http"

	_ "github.com/alifdwt/digitalent-golang/domain/responses/order"
	"github.com/alifdwt/digitalent-golang/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param body body models.Order true "Create order"
// @Success 200 {object} order.OrderResponse
// @Failure 400 {object} responses.ErrorMessage
// @Router /orders [post]
func (h *Handler) createOrder(ctx *gin.Context) {
	var body models.Order
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := h.services.Order.CreateOrder(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// @Summary Get all order
// @Description Get all order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} order.OrderResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /orders [get]
func (h *Handler) getAllOrder(ctx *gin.Context) {
	order, err := h.services.Order.GetAllOrder()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

type updateOrderURI struct {
	OrderID uint `uri:"orderId" binding:"required,min=1"`
}

// @Summary Update order
// @Description Update order
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Param body body models.Order true "Update order"
// @Success 200 {object} order.OrderResponse
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /orders/{orderId} [put]
func (h *Handler) updateOrder(ctx *gin.Context) {
	var uri updateOrderURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var body models.Order
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := h.services.Order.UpdateOrder(uri.OrderID, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// @Summary Delete order
// @Description Delete order
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path int true "Order ID"
// @Success 200 {object} string
// @Failure 400 {object} responses.ErrorMessage
// @Failure 500 {object} responses.ErrorMessage
// @Router /orders/{orderId} [delete]
func (h *Handler) deleteOrder(ctx *gin.Context) {
	var uri updateOrderURI
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := h.services.Order.DeleteOrder(uri.OrderID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Success delete")
}
