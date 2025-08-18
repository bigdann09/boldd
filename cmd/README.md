# BoldD Project

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
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

BoldD is a Go-based backend application designed to handle various functionalities such as user authentication, profile management, category and subcategory management, and more. It leverages modern libraries and frameworks like Gin for HTTP routing, GORM for database interactions, and Redis for caching.

---

## Features

- **User Authentication**: Includes user registration, OTP verification, and role assignment.
- **Profile Management**: Allows users to view and update their profiles, including password changes.
- **Category and Subcategory Management**: CRUD operations for categories and subcategories.
- **Caching**: Implements Redis caching for optimized performance.
- **Email Notifications**: Sends registration and OTP emails using pre-defined templates.
- **Database Migrations**: Supports database schema management using `migrate`.

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