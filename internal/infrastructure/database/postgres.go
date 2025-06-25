package database

import (
	"fmt"

	"prototype-fiber/internal/domain/entities"
	"prototype-fiber/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate tables
	if err := db.AutoMigrate(
		&entities.User{},
		&entities.Product{},
		&entities.Cart{},
		&entities.CartItem{},
		&entities.Order{},
		&entities.OrderItem{},
		&entities.Payment{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}