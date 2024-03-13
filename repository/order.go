package repository

import (
	"errors"
	"time"

	"github.com/alifdwt/digitalent-golang/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(request *models.Order) (*models.Order, error) {
	tx := r.db.Begin()

	orderCreate := models.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    time.Now(),
	}

	if err := tx.Create(&orderCreate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range request.Items {
		orderItem := models.Items{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     orderCreate.ID,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	checkOrder := tx.Preload("Items").Debug().Where("id = ?", orderCreate.ID).First(&orderCreate)
	if checkOrder.RowsAffected < 0 {
		return nil, errors.New("failed to get order")
	}

	tx.Commit()
	return &orderCreate, nil
}

func (r *orderRepository) GetAllOrder() (*[]models.Order, error) {
	var orders []models.Order

	db := r.db.Model(&orders)

	checkOrder := db.Preload("Items").Debug().Find(&orders)
	if checkOrder.RowsAffected < 1 {
		return nil, errors.New("there is no orders yet")
	}

	return &orders, nil
}

func (r *orderRepository) UpdateOrder(orderId uint, request *models.Order) (*models.Order, error) {
	var order models.Order

	tx := r.db.Begin()

	if err := tx.Preload("Items").Where("id = ?", orderId).First(&order).Error; err != nil {
		return nil, errors.New("failed to find product: " + err.Error())
	}

	order.CustomerName = request.CustomerName
	order.OrderedAt = request.OrderedAt

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to update order: " + err.Error())
	}

	for _, orderItem := range order.Items {
		if err := tx.Delete(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("failed to update order: " + err.Error())
		}
	}

	for _, item := range request.Items {
		orderItem := models.Items{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     order.ID,
		}

		// create new order item
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			return nil, errors.New("failed to update order: " + err.Error())
		}
	}

	checkOrder := tx.Preload("Items").Debug().Where("id = ?", orderId).First(&order)
	if checkOrder.RowsAffected < 0 {
		return nil, errors.New("failed to get order")
	}

	tx.Commit()
	return &order, nil
}

func (r *orderRepository) DeleteOrder(orderId uint) error {
	var order models.Order

	err := r.db.Preload("Items").Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return err
	}

	for _, orderItem := range order.Items {
		if err := r.db.Delete(&orderItem).Error; err != nil {
			return err
		}
	}

	return r.db.Delete(&order).Error
}
