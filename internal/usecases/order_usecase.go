package usecases

import (
	"errors"

	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
)

type OrderUseCase struct {
	orderRepo   entities.OrderRepository
	cartRepo    entities.CartRepository
	productRepo entities.ProductRepository
}

type CreateOrderRequest struct {
	ShippingAddress string `json:"shipping_address" validate:"required"`
	BillingAddress  string `json:"billing_address" validate:"required"`
}

func NewOrderUseCase(orderRepo entities.OrderRepository, cartRepo entities.CartRepository, productRepo entities.ProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (uc *OrderUseCase) CreateOrder(userID uuid.UUID, req *CreateOrderRequest) (*entities.Order, error) {
	cart, err := uc.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Check stock availability for all items
	for _, item := range cart.Items {
		product, err := uc.productRepo.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}
		if !product.CanFulfillQuantity(item.Quantity) {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}
	}

	// Create order
	order := &entities.Order{
		UserID:       userID,
		Status:       entities.OrderStatusPending,
		ShippingAddr: req.ShippingAddress,
		BillingAddr:  req.BillingAddress,
	}

	// Convert cart items to order items
	var orderItems []entities.OrderItem
	total := 0.0
	for _, cartItem := range cart.Items {
		orderItem := entities.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			Price:     cartItem.Price,
		}
		orderItems = append(orderItems, orderItem)
		total += cartItem.GetSubtotal()
	}

	order.Items = orderItems
	order.Total = total

	if err := uc.orderRepo.Create(order); err != nil {
		return nil, err
	}

	// Update product stock
	for _, item := range cart.Items {
		if err := uc.productRepo.UpdateStock(item.ProductID, -item.Quantity); err != nil {
			// Rollback order creation if stock update fails
			return nil, errors.New("failed to update stock")
		}
	}

	// Clear cart after successful order
	uc.cartRepo.Clear(cart.ID)

	return order, nil
}

func (uc *OrderUseCase) GetOrder(id uuid.UUID) (*entities.Order, error) {
	return uc.orderRepo.GetByID(id)
}

func (uc *OrderUseCase) GetUserOrders(userID uuid.UUID, offset, limit int) ([]*entities.Order, error) {
	return uc.orderRepo.GetByUserID(userID, offset, limit)
}

func (uc *OrderUseCase) UpdateOrderStatus(id uuid.UUID, status entities.OrderStatus) error {
	order, err := uc.orderRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Validate status transition
	if !uc.isValidStatusTransition(order.Status, status) {
		return errors.New("invalid status transition")
	}

	return uc.orderRepo.UpdateStatus(id, status)
}

func (uc *OrderUseCase) CancelOrder(userID, orderID uuid.UUID) error {
	order, err := uc.orderRepo.GetByID(orderID)
	if err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("unauthorized")
	}

	if !order.CanBeCancelled() {
		return errors.New("order cannot be cancelled")
	}

	// Restore stock
	for _, item := range order.Items {
		if err := uc.productRepo.UpdateStock(item.ProductID, item.Quantity); err != nil {
			return errors.New("failed to restore stock")
		}
	}

	return uc.orderRepo.UpdateStatus(orderID, entities.OrderStatusCancelled)
}

func (uc *OrderUseCase) ListOrders(offset, limit int) ([]*entities.Order, error) {
	return uc.orderRepo.List(offset, limit)
}

func (uc *OrderUseCase) isValidStatusTransition(current, new entities.OrderStatus) bool {
	validTransitions := map[entities.OrderStatus][]entities.OrderStatus{
		entities.OrderStatusPending:    {entities.OrderStatusPaid, entities.OrderStatusCancelled},
		entities.OrderStatusPaid:       {entities.OrderStatusProcessing, entities.OrderStatusCancelled, entities.OrderStatusRefunded},
		entities.OrderStatusProcessing: {entities.OrderStatusShipped, entities.OrderStatusCancelled},
		entities.OrderStatusShipped:    {entities.OrderStatusDelivered},
		entities.OrderStatusDelivered:  {entities.OrderStatusRefunded},
		entities.OrderStatusCancelled:  {},
		entities.OrderStatusRefunded:   {},
	}

	allowed, exists := validTransitions[current]
	if !exists {
		return false
	}

	for _, status := range allowed {
		if status == new {
			return true
		}
	}
	return false
}