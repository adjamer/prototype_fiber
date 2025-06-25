package handlers

import (
	"prototype-fiber/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewAuthHandler(userUseCase *usecases.UserUseCase) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUseCase,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req usecases.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	resp, err := h.userUseCase.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req usecases.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	resp, err := h.userUseCase.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(resp)
}