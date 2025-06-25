package tests

import (
	"testing"

	"prototype-fiber/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestProduct_IsInStock(t *testing.T) {
	// Product with stock and active
	product1 := &entities.Product{
		Stock:    10,
		IsActive: true,
	}
	assert.True(t, product1.IsInStock())

	// Product with no stock
	product2 := &entities.Product{
		Stock:    0,
		IsActive: true,
	}
	assert.False(t, product2.IsInStock())

	// Product inactive
	product3 := &entities.Product{
		Stock:    10,
		IsActive: false,
	}
	assert.False(t, product3.IsInStock())
}

func TestProduct_CanFulfillQuantity(t *testing.T) {
	product := &entities.Product{
		Stock:    10,
		IsActive: true,
	}

	// Can fulfill quantity within stock
	assert.True(t, product.CanFulfillQuantity(5))
	assert.True(t, product.CanFulfillQuantity(10))

	// Cannot fulfill quantity exceeding stock
	assert.False(t, product.CanFulfillQuantity(15))

	// Cannot fulfill if inactive
	product.IsActive = false
	assert.False(t, product.CanFulfillQuantity(5))
}