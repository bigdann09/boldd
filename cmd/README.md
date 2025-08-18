# Boldd Project

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
- [Environment Variables](#environment-variables)
- [Running the Application](#running-the-application)
- [Database Migrations](#database-migrations)
- [API Documentation](#api-documentation)
- [Caching](#caching)
- [Mail Templates](#mail-templates)
- [Contributing](#contributing)

---

## Introduction

BoldD is a multivendor eCommerce API designed to handle a wide range of functionalities required for a modern eCommerce platform. It supports features like user authentication, vendor management, product management, order processing, and more. The API is built using Go and leverages modern libraries and frameworks like Gin for HTTP routing, GORM for database interactions, and Redis for caching. The project is currently under active development.


---

## Features

### Core Features
- **User Authentication**: Includes user registration, OTP verification, login, and role-based access control.
- **Vendor Management**: Allows vendors to register, manage their profiles, and list products.
- **Product Management**: CRUD operations for products, categories, and subcategories.
- **Order Management**: Handles order placement, tracking, and status updates.
- **Cart Management**: Supports adding, updating, and removing items from the cart.
- **Transaction Management**: Manages payment processing and transaction history.
- **Invoice Generation**: Automatically generates invoices for completed orders.
- **Shipping and Delivery**: Tracks shipping details and delivery status.
- **Wishlist**: Allows users to save products for future purchases.
- **Review and Ratings**: Enables users to review and rate products.
- **Multivendor Support**: Allows multiple vendors to sell their products on the platform.

### Additional Features
- **Admin Dashboard**: Provides an interface for administrators to manage users, vendors, and products.
- **Notifications**: Sends email and in-app notifications for order updates, promotions, and more.
- **Analytics and Reporting**: Tracks sales, revenue, and other key metrics.
- **Discounts and Coupons**: Supports promotional discounts and coupon codes.
- **Search and Filters**: Implements advanced search and filtering options for products.
- **Inventory Management**: Tracks stock levels and notifies vendors of low inventory.
- **Multilingual and Multi-Currency Support**: Supports multiple languages and currencies for global reach.

---

## Project Structure

```plaintext
.
├── cmd/
│   └── api/
│       └── main.go          # Entry point for the application
├── internal/
│   ├── api/
│   │   ├── handlers/        # API handlers for various modules
│   │   ├── middlewares/     # Middleware for authentication and roles
│   │   └── routes/          # API route definitions
│   ├── application/
│   │   ├── attributes/      # Attribute-related business logic
│   │   ├── categories/      # Category-related business logic
│   │   ├── profile/         # Profile-related business logic
│   │   └── subcategories/   # Subcategory-related business logic
│   ├── config/              # Configuration loading and management
│   ├── domain/
│   │   ├── dtos/            # Data Transfer Objects
│   │   └── entities/        # Database entities
│   ├── infrastructure/
│   │   ├── cache/           # Redis caching implementation
│   │   ├── mail/            # Email templates and mailer logic
│   │   ├── persistence/     # Database and repository implementations
│   │   └── redis/           # Redis client setup
│   └── infrastructure/
│       └── persistence/
│           └── seeders/     # Database seeders
├── docker/
│   └── api/
│       └── Dockerfile       # Dockerfile for the API
├── Makefile                 # Makefile for common tasks
└── .env                     # Environment variables file
```

---

## Setup and Installation
Prerequisites

- Go 1.24 or later
- Docker and Docker Compose
- Postgres
- Redis

Installation Steps

1. Clone the repository:
```bash
git clone https://github.com/your-repo/boldd.git
cd boldd
```

2. Install dependencies:
```bash
go mod download
```

3. Set up the .env file: Copy .env.example file to .env and update the values as needed.

---

## Environment Variables
The application uses environment variables for configuration. Below are the key variables:

- Database Configuration:
    - DATABASE_HOST
    - DATABASE_PORT
    - DATABASE_USER
    - DATABASE_PASS
    - DATABASE_NAME
    - DATABASE_SSLMODE

- Redis Configuration:
    - REDIS_HOST
    - REDIS_DB
    - REDIS_HOST
    - REDIS_PORT
    - REDIS_PASSWORD

- Mail Configruation:
    - MAIL_HOST
    - MAIL_PORT
    - MAIL_FROM
    - MAIL_USERNAME
    -MAIL_PASSWORD

- Google OAuth Configuration:
    - GOOGLE_CLIENT_ID
    - GOOGLE_CLIENT_SECRET
    - GOOGLE_CALLBACK_URI

- Cloudinary Configuration:
    - CLOUDINARY_CLOUD_NAME
    - CLOUDINARY_API_KEY
    - CLOUDINARY_API_SECRET

- Grafana Configuration:
    - GRAFANA_USER
    - GRAFANA_PASSWORD

Refer to the `internal/config/config.go` file for more details.

---
## Configuring `config.yaml`
The `config.yaml` file is used to centralize application configuration. It is stored in the `$HOME/.config/boldd/config.yaml`

```yaml
application:
  port: 8003
  url: "URL"
  timezone: "TIMEZONE"
  environment: "ENVIRONMENT"
cors:
  allowed_origins: ["ORIGIN1", "ORIGIN2"]
jwt:
  key: "JWT KEY"
  access_expiry: EXPIRY IN HOUR (2 means 2 hours)
  refresh_expiry: EXPIRATION IN HOUR
```

Steps to Setup config.yaml
1. Navigate to the `$HOME/.config/boldd/` directory:
```bash
mkdir -p $HOME/.config/boldd
```

2. Create the `config.yaml` file:
```bash
mkdir -p $HOME/.config/boldd
```

1. Copy the example structure above into the file.
2. Update the values based on your environment (e.g., database credentials, Redis host, mail server settings).
3. Save and close the file.

The application will automatically load the config.yaml file from this location during startup. Ensure the file is properly formatted and contains all required fields.

---
### Running the Application
Local Development

1. Start the application
```bash
air
```

2. Access the application at `http://localhost:8003`.

---
## Database Migrations
Creating a New Migration
```bash
make migrate_create name=<migration_name>
```

Running the Migrations
```bash
make migrate_up
```

Rolling Back Migrations
```bash
make migrate_down
```

---
## API Documentation
The project uses Swagger for API documentation. To generate the documentation:

```bash
make swag-generate
```

Access the documentation at `http://localhost:8003/swagger/index.html`

---
## Caching
The application uses Redis for caching. The `Cache` utility in `internal/infrastructure/cache/redis.go` provides methods for:

- Setting and retrieving cached values
Setting and retrieving cached values
- Invalidating cache entries

---
## Mail Templates
The application includes pre-defined email templates for user registration and OTP verification. These templates are located in:

```bash
internal/infrastructure/mail/templates/
```

- `register.html`: Used for registration emails.
- `reset.html`: Used for password reset emails.
- `resend.html`: Used for OTP resend emails.

---
## Contributing
Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.