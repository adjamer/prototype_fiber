package entities

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID              uuid.UUID     `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	OrderID         uuid.UUID     `json:"order_id" gorm:"type:uuid;not null"`
	Amount          float64       `json:"amount" gorm:"not null"`
	Currency        string        `json:"currency" gorm:"default:'USD'"`
	Method          PaymentMethod `json:"method" gorm:"not null"`
	Status          PaymentStatus `json:"status" gorm:"default:'pending'"`
	ExternalID      string        `json:"external_id"`
	TransactionID   string        `json:"transaction_id"`
	ProcessorRef    string        `json:"processor_reference"`
	FailureReason   string        `json:"failure_reason"`
	ProcessedAt     *time.Time    `json:"processed_at"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type PaymentMethod string

const (
	PaymentMethodCard   PaymentMethod = "card"
	PaymentMethodPaypal PaymentMethod = "paypal"
	PaymentMethodBank   PaymentMethod = "bank_transfer"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

type PaymentRepository interface {
	Create(payment *Payment) error
	GetByID(id uuid.UUID) (*Payment, error)
	GetByOrderID(orderID uuid.UUID) (*Payment, error)
	Update(payment *Payment) error
	UpdateStatus(id uuid.UUID, status PaymentStatus) error
	List(offset, limit int) ([]*Payment, error)
}

func (p *Payment) IsSuccessful() bool {
	return p.Status == PaymentStatusCompleted
}

func (p *Payment) CanBeRefunded() bool {
	return p.Status == PaymentStatusCompleted
}