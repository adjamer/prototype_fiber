package usecases

import (
	"prototype-fiber/internal/domain/entities"
	"errors"

	"github.com/google/uuid"
)

type ProductUseCase struct {
	productRepo entities.ProductRepository
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"required"`
	Stock       int     `json:"stock" validate:"gte=0"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}

func NewProductUseCase(productRepo entities.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (uc *ProductUseCase) CreateProduct(req *CreateProductRequest) (*entities.Product, error) {
	// Check if SKU already exists
	existing, _ := uc.productRepo.GetBySKU(req.SKU)
	if existing != nil {
		return nil, errors.New("product with this SKU already exists")
	}

	product := &entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		SKU:         req.SKU,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
		IsActive:    true,
	}

	if err := uc.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUseCase) GetProduct(id uuid.UUID) (*entities.Product, error) {
	return uc.productRepo.GetByID(id)
}

func (uc *ProductUseCase) ListProducts(offset, limit int, category string) ([]*entities.Product, error) {
	return uc.productRepo.List(offset, limit, category)
}

func (uc *ProductUseCase) SearchProducts(query string, offset, limit int) ([]*entities.Product, error) {
	return uc.productRepo.Search(query, offset, limit)
}

func (uc *ProductUseCase) UpdateProduct(id uuid.UUID, updates map[string]interface{}) (*entities.Product, error) {
	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if name, ok := updates["name"].(string); ok {
		product.Name = name
	}
	if description, ok := updates["description"].(string); ok {
		product.Description = description
	}
	if price, ok := updates["price"].(float64); ok {
		product.Price = price
	}
	if stock, ok := updates["stock"].(int); ok {
		product.Stock = stock
	}
	if category, ok := updates["category"].(string); ok {
		product.Category = category
	}
	if imageURL, ok := updates["image_url"].(string); ok {
		product.ImageURL = imageURL
	}
	if isActive, ok := updates["is_active"].(bool); ok {
		product.IsActive = isActive
	}

	if err := uc.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUseCase) DeleteProduct(id uuid.UUID) error {
	return uc.productRepo.Delete(id)
}

func (uc *ProductUseCase) UpdateStock(id uuid.UUID, quantity int) error {
	return uc.productRepo.UpdateStock(id, quantity)
}