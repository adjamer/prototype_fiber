package main

import (
	"log"

	"prototype-fiber/internal/infrastructure/cache"
	"prototype-fiber/internal/infrastructure/database"
	"prototype-fiber/internal/interfaces/http/handlers"
	"prototype-fiber/internal/interfaces/http/routes"
	"prototype-fiber/internal/interfaces/repositories"
	"prototype-fiber/internal/usecases"
	"prototype-fiber/pkg/config"
	"prototype-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "prototype-fiber/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for the Fiber API.
func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger := logger.New()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	logger.Info("Connected to PostgreSQL database")

	// Initialize Redis
	redisClient, err := cache.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Warn("Failed to connect to Redis:", err)
	} else {
		_ = redisClient
		logger.Info("Connected to Redis")
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo, cfg.JWT)
	productUseCase := usecases.NewProductUseCase(productRepo)
	cartUseCase := usecases.NewCartUseCase(cartRepo, productRepo)
	orderUseCase := usecases.NewOrderUseCase(orderRepo, cartRepo, productRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userUseCase)
	userHandler := handlers.NewUserHandler(userUseCase)
	productHandler := handlers.NewProductHandler(productUseCase)
	cartHandler := handlers.NewCartHandler(cartUseCase)
	orderHandler := handlers.NewOrderHandler(orderUseCase)

	handlersStruct := &routes.Handlers{
		Auth:    authHandler,
		User:    userHandler,
		Product: productHandler,
		Cart:    cartHandler,
		Order:   orderHandler,
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Setup routes
	routes.SetupRoutes(app, handlersStruct, cfg.JWT.Secret)

	// Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app":    cfg.App.Name,
		})
	})

	// Start server
	logger.Infof("Starting %s on port %s", cfg.App.Name, cfg.App.Port)
	if err := app.Listen(":" + cfg.App.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}