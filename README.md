# Go Gin REST API Boilerplate
A production-ready boilerplate for building RESTful APIs in Go using Gin, PostgreSQL, and JWT-based authentication. This boilerplate follows clean architecture principles and provides a modular structure that is ready for extension.

## Features

- 🔐 JWT-based authentication
- 🏗️ Clean Architecture
- 📦 Modular structure
- 🗄️ PostgreSQL database integration
- 🔒 Password hashing

TODO: 

- 🧪 Testing setup
- ✅ Input validation
- 🚀 Production-ready
- 📝 API documentation

## Requirements

- [Go 1.21](https://golang.org/dl/) or higher
- [PostgreSQL 12](https://www.postgresql.org/download/) or higher
- `Make` (optional, for using Make commands)

## Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/GuiFernandess/go-gin-api-boilerplate.git
   cd go-gin-api-boilerplate

## Project Structure

api/
├── cmd/            # Application entry points
├── config/         # Configuration settings
├── internal/       # Private application code
│   ├── auth/       # Authentication logic (login, signup, token handling)
│   ├── routes/     # Routes and route handling
│   └── utils/      # Helper functions and utilities
└── migrations/     # Database migration files