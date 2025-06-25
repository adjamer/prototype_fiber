package usecases

import (
	"errors"

	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
)

type CartUseCase struct {
	cartRepo    entities.CartRepository
	productRepo entities.ProductRepository
}

type AddToCartRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
}

func NewCartUseCase(cartRepo entities.CartRepository, productRepo entities.ProductRepository) *CartUseCase {
	return &CartUseCase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (uc *CartUseCase) GetOrCreateCart(userID uuid.UUID) (*entities.Cart, error) {
	cart, err := uc.cartRepo.GetByUserID(userID)
	if err != nil {
		// Create new cart if doesn't exist
		cart = &entities.Cart{
			UserID: userID,
		}
		if err := uc.cartRepo.Create(cart); err != nil {
			return nil, err
		}
	}
	return cart, nil
}

func (uc *CartUseCase) AddToCart(userID uuid.UUID, req *AddToCartRequest) (*entities.Cart, error) {
	cart, err := uc.GetOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	product, err := uc.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if !product.CanFulfillQuantity(req.Quantity) {
		return nil, errors.New("insufficient stock")
	}

	cartItem := &entities.CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Price:     product.Price,
	}

	if err := uc.cartRepo.AddItem(cart.ID, cartItem); err != nil {
		return nil, err
	}

	return uc.cartRepo.GetByID(cart.ID)
}

func (uc *CartUseCase) UpdateCartItem(userID, productID uuid.UUID, quantity int) (*entities.Cart, error) {
	cart, err := uc.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if quantity <= 0 {
		if err := uc.cartRepo.RemoveItem(cart.ID, productID); err != nil {
			return nil, err
		}
	} else {
		product, err := uc.productRepo.GetByID(productID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		if !product.CanFulfillQuantity(quantity) {
			return nil, errors.New("insufficient stock")
		}

		if err := uc.cartRepo.UpdateItem(cart.ID, productID, quantity); err != nil {
			return nil, err
		}
	}

	return uc.cartRepo.GetByID(cart.ID)
}

func (uc *CartUseCase) RemoveFromCart(userID, productID uuid.UUID) (*entities.Cart, error) {
	cart, err := uc.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	if err := uc.cartRepo.RemoveItem(cart.ID, productID); err != nil {
		return nil, err
	}

	return uc.cartRepo.GetByID(cart.ID)
}

func (uc *CartUseCase) ClearCart(userID uuid.UUID) error {
	cart, err := uc.cartRepo.GetByUserID(userID)
	if err != nil {
		return err
	}

	return uc.cartRepo.Clear(cart.ID)
}

func (uc *CartUseCase) GetCart(userID uuid.UUID) (*entities.Cart, error) {
	return uc.cartRepo.GetByUserID(userID)
}