package tests

import (
	"prototype-fiber/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
	user := &entities.User{}
	password := "testpassword123"

	err := user.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, password, user.Password)
}

func TestUser_CheckPassword(t *testing.T) {
	user := &entities.User{}
	password := "testpassword123"

	// Hash the password
	err := user.HashPassword(password)
	assert.NoError(t, err)

	// Check correct password
	assert.True(t, user.CheckPassword(password))

	// Check incorrect password
	assert.False(t, user.CheckPassword("wrongpassword"))
}

func TestUser_FullName(t *testing.T) {
	user := &entities.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	assert.Equal(t, "John Doe", user.FullName())
}