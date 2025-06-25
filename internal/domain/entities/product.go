package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	SKU         string    `json:"sku" gorm:"uniqueIndex;not null"`
	Stock       int       `json:"stock" gorm:"default:0"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"image_url"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id uuid.UUID) (*Product, error)
	GetBySKU(sku string) (*Product, error)
	Update(product *Product) error
	Delete(id uuid.UUID) error
	List(offset, limit int, category string) ([]*Product, error)
	Search(query string, offset, limit int) ([]*Product, error)
	UpdateStock(id uuid.UUID, quantity int) error
}

func (p *Product) IsInStock() bool {
	return p.Stock > 0 && p.IsActive
}

func (p *Product) CanFulfillQuantity(quantity int) bool {
	return p.Stock >= quantity && p.IsActive
}