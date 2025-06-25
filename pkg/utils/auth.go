package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUserIDFromContext(c *fiber.Ctx) (uuid.UUID, error) {
	userIDStr, ok := c.Locals("user_id").(string)
	if !ok {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, errors.New("invalid user ID format")
	}

	return userID, nil
}

func GetUserRoleFromContext(c *fiber.Ctx) (string, error) {
	role, ok := c.Locals("role").(string)
	if !ok {
		return "", errors.New("user role not found in context")
	}
	return role, nil
}