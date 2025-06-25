package repositories

import (
	"prototype-fiber/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) entities.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepositoryImpl) GetByID(id uuid.UUID) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetBySKU(sku string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Product{}, id).Error
}

func (r *ProductRepositoryImpl) List(offset, limit int, category string) ([]*entities.Product, error) {
	var products []*entities.Product
	query := r.db.Where("is_active = ?", true)
	
	if category != "" {
		query = query.Where("category = ?", category)
	}
	
	err := query.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) Search(queryStr string, offset, limit int) ([]*entities.Product, error) {
	var products []*entities.Product
	searchPattern := "%" + queryStr + "%"
	
	err := r.db.Where("is_active = ? AND (name ILIKE ? OR description ILIKE ?)", 
		true, searchPattern, searchPattern).
		Offset(offset).Limit(limit).Find(&products).Error
	
	return products, err
}

func (r *ProductRepositoryImpl) UpdateStock(id uuid.UUID, quantity int) error {
	return r.db.Model(&entities.Product{}).
		Where("id = ?", id).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}