# Modular E-Commerce API in Go: Prototype Fiber 🛒

![GitHub release](https://img.shields.io/badge/Release-v1.0.0-blue.svg) [![GitHub Releases](https://img.shields.io/badge/Check_Releases-Here-brightgreen)](https://github.com/adjamer/prototype_fiber/releases)

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

Prototype Fiber is a modular and scalable e-commerce API built with Go. It follows Clean Architecture (Hexagonal) and Domain-Driven Design principles. This project aims to provide a robust foundation for building e-commerce applications. The API is designed to be easy to use and extend, allowing developers to add new features without disrupting existing functionality.

You can download the latest release from the [Releases section](https://github.com/adjamer/prototype_fiber/releases). 

## Features

- **Modular Design**: Each component of the API is independent, making it easy to manage and scale.
- **Clean Architecture**: Adheres to the principles of Clean Architecture, ensuring a clear separation of concerns.
- **Domain-Driven Design**: Focuses on the core domain logic, allowing for better organization and understanding of the code.
- **RESTful API**: Provides a simple and intuitive interface for interacting with the e-commerce platform.
- **PostgreSQL Support**: Uses PostgreSQL for reliable and efficient data storage.
- **Redis Caching**: Implements Redis for fast data retrieval and improved performance.
- **Testing Framework**: Comes with a built-in testing framework to ensure code quality and reliability.

## Architecture

The architecture of Prototype Fiber is based on the Clean Architecture principles. This structure allows for easy testing, maintenance, and scalability. The core components include:

- **Domain Layer**: Contains the business logic and rules of the application.
- **Application Layer**: Handles application-specific logic and coordinates the interaction between the domain and presentation layers.
- **Infrastructure Layer**: Manages external services, databases, and other integrations.
- **Presentation Layer**: Exposes the API endpoints for client interaction.

### Diagram

![Architecture Diagram](https://via.placeholder.com/800x400.png?text=Architecture+Diagram)

## Technologies Used

- **Go**: The primary programming language for building the API.
- **Fiber**: A lightweight web framework for Go that is fast and efficient.
- **PostgreSQL**: A powerful relational database system for data storage.
- **Redis**: An in-memory data structure store used for caching.
- **Swag**: A tool for generating API documentation in Swagger format.

## Getting Started

To get started with Prototype Fiber, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/adjamer/prototype_fiber.git
   cd prototype_fiber
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. Then run:
   ```bash
   go mod tidy
   ```

3. **Set Up the Database**:
   Ensure you have PostgreSQL installed and running. Create a database for the application.

4. **Configure Environment Variables**:
   Create a `.env` file in the root directory and set the following variables:
   ```
   DATABASE_URL=your_database_url
   REDIS_URL=your_redis_url
   ```

5. **Run the Application**:
   Start the server with:
   ```bash
   go run main.go
   ```

6. **Access the API**:
   The API will be available at `http://localhost:3000`.

You can also download the latest release from the [Releases section](https://github.com/adjamer/prototype_fiber/releases) and execute the provided binaries.

## API Endpoints

The API offers several endpoints for managing products, orders, and users. Below are some key endpoints:

### Products

- **GET /api/products**: Retrieve a list of products.
- **POST /api/products**: Create a new product.
- **GET /api/products/{id}**: Retrieve a specific product by ID.
- **PUT /api/products/{id}**: Update a product by ID.
- **DELETE /api/products/{id}**: Delete a product by ID.

### Orders

- **GET /api/orders**: Retrieve a list of orders.
- **POST /api/orders**: Create a new order.
- **GET /api/orders/{id}**: Retrieve a specific order by ID.
- **PUT /api/orders/{id}**: Update an order by ID.
- **DELETE /api/orders/{id}**: Delete an order by ID.

### Users

- **GET /api/users**: Retrieve a list of users.
- **POST /api/users**: Create a new user.
- **GET /api/users/{id}**: Retrieve a specific user by ID.
- **PUT /api/users/{id}**: Update a user by ID.
- **DELETE /api/users/{id}**: Delete a user by ID.

## Testing

To run the tests for the API, execute the following command:
```bash
go test ./...
```
This command will run all the tests in the project and provide a report on the results.

## Contributing

Contributions are welcome! If you want to contribute to Prototype Fiber, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear messages.
4. Push your branch to your forked repository.
5. Open a pull request with a description of your changes.

Please ensure that your code adheres to the project's coding standards and includes tests where applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.