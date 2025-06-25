package routes

import (
	"prototype-fiber/internal/interfaces/http/handlers"
	"prototype-fiber/internal/interfaces/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Auth    *handlers.AuthHandler
	User    *handlers.UserHandler
	Product *handlers.ProductHandler
	Cart    *handlers.CartHandler
	Order   *handlers.OrderHandler
}

func SetupRoutes(app *fiber.App, handlers *Handlers, jwtSecret string) {
	api := app.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", handlers.Auth.Register)
	auth.Post("/login", handlers.Auth.Login)

	// Product routes (public for reading, admin for writing)
	products := api.Group("/products")
	products.Get("/", handlers.Product.ListProducts)
	products.Get("/search", handlers.Product.SearchProducts)
	products.Get("/:id", handlers.Product.GetProduct)

	// Protected routes
	protected := api.Use(middleware.AuthMiddleware(jwtSecret))

	// User routes
	users := protected.Group("/users")
	users.Get("/profile", handlers.User.GetProfile)
	users.Put("/profile", handlers.User.UpdateProfile)

	// Cart routes
	cart := protected.Group("/cart")
	cart.Get("/", handlers.Cart.GetCart)
	cart.Post("/items", handlers.Cart.AddToCart)
	cart.Put("/items/:productId", handlers.Cart.UpdateCartItem)
	cart.Delete("/items/:productId", handlers.Cart.RemoveFromCart)
	cart.Delete("/", handlers.Cart.ClearCart)

	// Order routes
	orders := protected.Group("/orders")
	orders.Post("/", handlers.Order.CreateOrder)
	orders.Get("/", handlers.Order.GetUserOrders)
	orders.Get("/:id", handlers.Order.GetOrder)
	orders.Delete("/:id", handlers.Order.CancelOrder)

	// Admin routes
	admin := protected.Use(middleware.AdminMiddleware())
	adminProducts := admin.Group("/admin/products")
	adminProducts.Post("/", handlers.Product.CreateProduct)
	adminProducts.Put("/:id", handlers.Product.UpdateProduct)
	adminProducts.Delete("/:id", handlers.Product.DeleteProduct)
}