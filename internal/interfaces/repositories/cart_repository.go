package repositories

import (
	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) entities.CartRepository {
	return &CartRepositoryImpl{db: db}
}

func (r *CartRepositoryImpl) Create(cart *entities.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepositoryImpl) GetByUserID(userID uuid.UUID) (*entities.Cart, error) {
	var cart entities.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepositoryImpl) GetByID(id uuid.UUID) (*entities.Cart, error) {
	var cart entities.Cart
	err := r.db.Preload("Items.Product").Where("id = ?", id).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepositoryImpl) Update(cart *entities.Cart) error {
	return r.db.Save(cart).Error
}

func (r *CartRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Cart{}, id).Error
}

func (r *CartRepositoryImpl) AddItem(cartID uuid.UUID, item *entities.CartItem) error {
	item.CartID = cartID
	
	// Check if item already exists
	var existingItem entities.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, item.ProductID).First(&existingItem).Error
	
	if err == gorm.ErrRecordNotFound {
		// Create new item
		return r.db.Create(item).Error
	} else if err != nil {
		return err
	} else {
		// Update existing item quantity
		existingItem.Quantity += item.Quantity
		return r.db.Save(&existingItem).Error
	}
}

func (r *CartRepositoryImpl) UpdateItem(cartID uuid.UUID, productID uuid.UUID, quantity int) error {
	return r.db.Model(&entities.CartItem{}).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		Update("quantity", quantity).Error
}

func (r *CartRepositoryImpl) RemoveItem(cartID uuid.UUID, productID uuid.UUID) error {
	return r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).
		Delete(&entities.CartItem{}).Error
}

func (r *CartRepositoryImpl) Clear(cartID uuid.UUID) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&entities.CartItem{}).Error
}