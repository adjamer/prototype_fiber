# Prototype Fiber

A modular and scalable e-commerce API built with Go, following Clean Architecture (Hexagonal) and Domain-Driven Design principles.

## Features

- **Clean Architecture**: Well-organized code with clear separation of concerns
- **Domain-Driven Design**: Business logic encapsulated in domain entities
- **RESTful API**: Comprehensive endpoints for e-commerce operations
- **Authentication & Authorization**: JWT-based auth with role-based access
- **Database Integration**: PostgreSQL with GORM ORM
- **Caching**: Redis for improved performance
- **Docker Support**: Easy deployment with Docker Compose
- **Unit Testing**: Comprehensive test coverage
- **API Documentation**: Well-documented endpoints

## Architecture

```bash
cmd/                    # Application entry points
├── api/                # Main API server
internal/               # Internal application code
├── domain/             # Domain layer (entities, interfaces)
│   ├── entities/       # Business entities
│   └── repositories/   # Repository interfaces
├── usecases/           # Application business logic
├── infrastructure/     # External concerns (database, cache)
├── interfaces/         # Interface adapters
│   ├── http/           # HTTP handlers and routes
│   └── repositories/   # Repository implementations
pkg/                    # Shared packages
├── config/             # Configuration management
├── logger/             # Logging utilities
└── utils/              # Common utilities
```

## Modules

### Users

- User registration and authentication
- Profile management
- Role-based access control (Customer, Admin)

### Products

- Product catalog management
- Search and filtering
- Stock management
- Category organization

### Shopping Cart

- Add/remove items
- Update quantities
- Cart persistence per user
- Price calculations

### Orders

- Order creation from cart
- Order status tracking
- Order history
- Cancellation support

### Payments

- Payment processing integration ready
- Multiple payment methods support
- Transaction tracking

## API Endpoints

### Authentication Endpoints

- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login

### Users Endpoints

- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update user profile

### Products Endpoints

- `GET /api/v1/products` - List products
- `GET /api/v1/products/:id` - Get product details
- `GET /api/v1/products/search` - Search products
- `POST /api/v1/admin/products` - Create product (Admin)
- `PUT /api/v1/admin/products/:id` - Update product (Admin)
- `DELETE /api/v1/admin/products/:id` - Delete product (Admin)

### Cart Endpoints

- `GET /api/v1/cart` - Get user cart
- `POST /api/v1/cart/items` - Add item to cart
- `PUT /api/v1/cart/items/:productId` - Update cart item
- `DELETE /api/v1/cart/items/:productId` - Remove item from cart
- `DELETE /api/v1/cart` - Clear cart

### Orders Endpoints

- `POST /api/v1/orders` - Create order
- `GET /api/v1/orders` - Get user orders
- `GET /api/v1/orders/:id` - Get order details
- `DELETE /api/v1/orders/:id` - Cancel order

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 13+
- Redis 6+
- Docker & Docker Compose (optional)

### Environment Setup

1. Copy the environment file:

```bash
cp .env.example .env
```

2. Update the `.env` file with your configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=ecommerce
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=your-secret-key
```

### Running with Docker Compose (Recommended)

```bash
# Start all services
make docker-up

# Build and start with logs
make docker-build

# View logs
make logs

# Stop services
make docker-down
```

### Running Locally

1. Start PostgreSQL and Redis services

2. Install dependencies:

```bash
make deps
```

3. Run the application:

```bash
make run
```

### Running Tests

```bash
make test
```

## Usage Examples

### Register a new user

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```

### Create a product (Admin)

```bash
curl -X POST http://localhost:8080/api/v1/admin/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Laptop",
    "description": "High-performance laptop",
    "price": 999.99,
    "sku": "LAP001",
    "stock": 50,
    "category": "Electronics"
  }'
```

### Add item to cart

```bash
curl -X POST http://localhost:8080/api/v1/cart/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "product_id": "PRODUCT_UUID",
    "quantity": 2
  }'
```

## Development

### Code Structure

- **Domain Layer**: Contains business entities and rules
- **Use Case Layer**: Orchestrates business logic
- **Interface Layer**: Handles external communication
- **Infrastructure Layer**: Implements external dependencies

### Adding New Features

1. Define entities in `internal/domain/entities/`
2. Create repository interfaces in domain
3. Implement use cases in `internal/usecases/`
4. Add repository implementations in `internal/interfaces/repositories/`
5. Create HTTP handlers in `internal/interfaces/http/handlers/`
6. Update routes in `internal/interfaces/http/routes/`

### Testing Strategy

- Unit tests for domain entities
- Integration tests for repositories
- Handler tests for HTTP endpoints
- Use dependency injection for easy mocking

## Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

This project is licensed under the MIT License.
