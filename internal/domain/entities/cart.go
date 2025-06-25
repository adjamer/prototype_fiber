package entities

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	User      User       `json:"user" gorm:"foreignKey:UserID"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CartID    uuid.UUID `json:"cart_id" gorm:"type:uuid;not null"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid;not null"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartRepository interface {
	Create(cart *Cart) error
	GetByUserID(userID uuid.UUID) (*Cart, error)
	GetByID(id uuid.UUID) (*Cart, error)
	Update(cart *Cart) error
	Delete(id uuid.UUID) error
	AddItem(cartID uuid.UUID, item *CartItem) error
	UpdateItem(cartID uuid.UUID, productID uuid.UUID, quantity int) error
	RemoveItem(cartID uuid.UUID, productID uuid.UUID) error
	Clear(cartID uuid.UUID) error
}

func (c *Cart) GetTotal() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.GetSubtotal()
	}
	return total
}

func (c *Cart) GetItemCount() int {
	count := 0
	for _, item := range c.Items {
		count += item.Quantity
	}
	return count
}

func (ci *CartItem) GetSubtotal() float64 {
	return ci.Price * float64(ci.Quantity)
}