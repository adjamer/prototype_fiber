package handlers

import (
	"prototype-fiber/internal/usecases"
	"prototype-fiber/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CartHandler struct {
	cartUseCase *usecases.CartUseCase
}

func NewCartHandler(cartUseCase *usecases.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUseCase: cartUseCase,
	}
}

func (h *CartHandler) GetCart(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	cart, err := h.cartUseCase.GetCart(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cart not found",
		})
	}

	return c.JSON(cart)
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var req usecases.AddToCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cart, err := h.cartUseCase.AddToCart(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cart)
}

func (h *CartHandler) UpdateCartItem(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	productID, err := uuid.Parse(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	var req struct {
		Quantity int `json:"quantity" validate:"required,gte=0"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cart, err := h.cartUseCase.UpdateCartItem(userID, productID, req.Quantity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cart)
}

func (h *CartHandler) RemoveFromCart(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	productID, err := uuid.Parse(c.Params("productId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid product ID",
		})
	}

	cart, err := h.cartUseCase.RemoveFromCart(userID, productID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cart)
}

func (h *CartHandler) ClearCart(c *fiber.Ctx) error {
	userID, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	if err := h.cartUseCase.ClearCart(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to clear cart",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}