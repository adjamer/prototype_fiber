package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID   `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        uuid.UUID   `json:"user_id" gorm:"type:uuid;not null"`
	User          User        `json:"user" gorm:"foreignKey:UserID"`
	Items         []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Status        OrderStatus `json:"status" gorm:"default:'pending'"`
	Total         float64     `json:"total" gorm:"not null"`
	ShippingCost  float64     `json:"shipping_cost" gorm:"default:0"`
	Tax           float64     `json:"tax" gorm:"default:0"`
	PaymentID     *uuid.UUID  `json:"payment_id,omitempty" gorm:"type:uuid"`
	Payment       *Payment    `json:"payment,omitempty" gorm:"foreignKey:PaymentID"`
	ShippingAddr  string      `json:"shipping_address"`
	BillingAddr   string      `json:"billing_address"`
	TrackingCode  string      `json:"tracking_code"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrderID   uuid.UUID `json:"order_id" gorm:"type:uuid;not null"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id uuid.UUID) (*Order, error)
	GetByUserID(userID uuid.UUID, offset, limit int) ([]*Order, error)
	Update(order *Order) error
	UpdateStatus(id uuid.UUID, status OrderStatus) error
	List(offset, limit int) ([]*Order, error)
	GetByStatus(status OrderStatus, offset, limit int) ([]*Order, error)
}

func (o *Order) CanBeCancelled() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusPaid
}

func (o *Order) CanBeRefunded() bool {
	return o.Status == OrderStatusPaid || o.Status == OrderStatusProcessing || o.Status == OrderStatusShipped
}

func (oi *OrderItem) GetSubtotal() float64 {
	return oi.Price * float64(oi.Quantity)
}