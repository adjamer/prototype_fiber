package repositories

import (
	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) entities.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(order *entities.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepositoryImpl) GetByID(id uuid.UUID) (*entities.Order, error) {
	var order entities.Order
	err := r.db.Preload("Items.Product").Preload("Payment").Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepositoryImpl) GetByUserID(userID uuid.UUID, offset, limit int) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := r.db.Preload("Items.Product").Preload("Payment").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).Limit(limit).Find(&orders).Error
	return orders, err
}

func (r *OrderRepositoryImpl) Update(order *entities.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepositoryImpl) UpdateStatus(id uuid.UUID, status entities.OrderStatus) error {
	return r.db.Model(&entities.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (r *OrderRepositoryImpl) List(offset, limit int) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := r.db.Preload("Items.Product").Preload("Payment").
		Order("created_at DESC").
		Offset(offset).Limit(limit).Find(&orders).Error
	return orders, err
}

func (r *OrderRepositoryImpl) GetByStatus(status entities.OrderStatus, offset, limit int) ([]*entities.Order, error) {
	var orders []*entities.Order
	err := r.db.Preload("Items.Product").Preload("Payment").
		Where("status = ?", status).
		Order("created_at DESC").
		Offset(offset).Limit(limit).Find(&orders).Error
	return orders, err
}