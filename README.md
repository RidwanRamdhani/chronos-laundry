# Chronos Laundry 🧺

[![Go Version](https://img.shields.io/badge/Go-1.25.4-00ADD8?logo=go)](https://golang.org/)
[![Vite](https://img.shields.io/badge/Vite-7.2.4-646CFF?logo=vite)](https://vitejs.dev/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

> A modern, full-stack laundry management system built with Go and Vite for efficient business operations, real-time tracking, and seamless customer experience.

## 📖 Overview

Chronos Laundry is a comprehensive laundry management solution designed to streamline operations for laundry businesses of all sizes. Built with a robust Go backend and a responsive Vite-powered frontend, it provides real-time transaction tracking, service management, and administrative controls.

### Key Highlights

- 🚀 **High Performance** - Built with Go for fast, concurrent processing
- 🎨 **Modern UI** - Responsive interface powered by Vite
- 🔐 **Secure Authentication** - JWT-based authentication system
- 📊 **Real-time Tracking** - Monitor transactions and status updates instantly
- 💰 **Dynamic Pricing** - Flexible service pricing management

## ✨ Features

### Core Functionality

- **Transaction Management**
  - Create, update, and track laundry transactions
  - Automatic transaction code generation
  - Multi-item transaction support
  - Transaction history tracking

- **Service Price Management**
  - Dynamic service pricing configuration
  - Easy price updates and management

- **Authentication & Authorization**
  - Secure JWT-based authentication
  - Role-based access control
  - Password encryption with bcrypt

- **Real-time Dashboard**
  - Transaction overview and statistics

- **Customer Tracking**
  - Track laundry status by transaction code
  - Real-time status updates
  - Transparent process visibility

## 🏗️ Architecture

### Technology Stack

#### Backend
- **Language**: Go 1.25.4
- **Framework**: Gin Web Framework
- **Database**: MySQL with GORM ORM
- **Authentication**: JWT (golang-jwt/jwt)
- **Security**: bcrypt password hashing

#### Frontend
- **Build Tool**: Vite 7.2.4
- **Architecture**: Vanilla JavaScript (modular)
- **Styling**: Custom CSS
- **HTTP Client**: Fetch API

### Project Structure

```
chronos-laundry/
├── backend/                    # Go backend application
│   ├── cmd/
│   │   └── main.go            # Application entry point
│   ├── config/
│   │   └── database.go        # Database configuration
│   ├── controllers/           # HTTP request handlers
│   │   ├── auth_controller.go
│   │   ├── service_price_controller.go
│   │   └── transaction_controller.go
│   ├── middlewares/           # HTTP middlewares
│   │   └── auth_middleware.go
│   ├── models/                # Data models
│   │   ├── admin.go
│   │   ├── service_price.go
│   │   ├── transaction.go
│   │   ├── transaction_item.go
│   │   └── transaction_history.go
│   ├── repositories/          # Data access layer
│   │   ├── admin_repository.go
│   │   ├── service_price_repository.go
│   │   ├── transaction_repository.go
│   │   └── transaction_history_repository.go
│   ├── routes/                # Route definitions
│   │   ├── router.go
│   │   ├── auth_routes.go
│   │   ├── service_price_routes.go
│   │   └── transaction_routes.go
│   ├── services/              # Business logic layer
│   │   ├── auth_service.go
│   │   ├── service_price_service.go
│   │   └── transaction_service.go
│   ├── utils/                 # Utility functions
│   │   ├── jwt.go
│   │   ├── password.go
│   │   ├── response.go
│   │   └── transcation_code.go
│   ├── go.mod                 # Go module definition
│   └── go.sum                 # Go dependencies checksum
│
├── frontend/frontend/         # Vite frontend application
│   ├── pages/                 # HTML pages
│   │   ├── dashboard.html
│   │   ├── login.html
│   │   ├── transactions.html
│   │   ├── transaction-create.html
│   │   ├── transaction-detail.html
│   │   ├── transaction-update.html
│   │   ├── service-prices.html
│   │   ├── service-price-create.html
│   │   ├── service-price-update.html
│   │   └── tracking.html
│   ├── js/                    # JavaScript modules
│   │   ├── dashboard.js
│   │   ├── login.js
│   │   ├── transactions.js
│   │   ├── transaction-create.js
│   │   ├── transaction-detail.js
│   │   ├── transaction-update.js
│   │   ├── service-prices.js
│   │   ├── service-price-create.js
│   │   ├── service-price-update.js
│   │   ├── tracking.js
│   │   └── layout.js
│   ├── layouts/
│   │   └── layout.html        # Shared layout template
│   ├── src/
│   │   ├── main.js            # Application entry point
│   │   ├── style.css          # Global styles
│   │   └── layout.css         # Layout styles
│   ├── public/                # Static assets
│   │   └── images/
│   ├── package.json           # Node dependencies
│   └── index.html             # Main HTML file
│
├── .gitignore                 # Git ignore rules
└── README.md                  # This file
```

## 🚀 Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.25.4 or higher ([Download](https://golang.org/dl/))
- **Node.js** 18.x or higher ([Download](https://nodejs.org/))
- **MySQL** 8.0 or higher ([Download](https://dev.mysql.com/downloads/))
- **Git** ([Download](https://git-scm.com/downloads))

### Installation

#### 1. Clone the Repository

```bash
git clone https://github.com/RidwanRamdhani/chronos-laundry.git
cd chronos-laundry
```

#### 2. Backend Setup

```bash
# Navigate to backend directory
cd backend

# Install Go dependencies
go mod download

# Create .env file
cp .env.example .env

# Configure your .env file with database credentials
# Example:
# DB_HOST=localhost
# DB_PORT=3306
# DB_USER=root
# DB_PASSWORD=your_password
# DB_NAME=chronos_laundry
# JWT_SECRET=your_jwt_secret_key
```

#### 3. Database Setup

```bash
# Create database
mysql -u root -p
CREATE DATABASE chronos_laundry;
exit;

# Run migrations (automatic on first run)
# The application will auto-migrate tables on startup
```

#### 4. Database Seeding

The project includes seeders to populate initial data for development and testing.

**Available Seeders:**

1. **Admin Seeder** - Creates default admin account
2. **Service Price Seeder** - Populates service prices for laundry items

**Running Seeders:**

```bash
# Navigate to backend directory
cd backend

# Run Admin Seeder
go run cmd/seeder/admin_seeder/admin.go

# Run Service Price Seeder
go run cmd/seeder/service_price_seeder/service_prices.go
```

**Default Admin Credentials:**
- **Username**: `admin`
- **Password**: `admin123`
- **Email**: `admin@chronos-laundry.com`
- **Full Name**: System Administrator

⚠️ **Important**: To use custom admin credentials, edit the seeder file at [`backend/cmd/seeder/admin_seeder/admin.go`](backend/cmd/seeder/admin_seeder/admin.go:28) before running the seeder. Modify the username, password, email, and full name values as needed.

**Service Price Categories:**

💡 **Tip**: To customize service prices, edit the seeder file at [`backend/cmd/seeder/service_price_seeder/service_prices.go`](backend/cmd/seeder/service_price_seeder/service_prices.go:25) before running the seeder. You can modify prices, add new items, or change service types as needed.

The service price seeder includes:
- **Regular Service** (reguler)
  - Cuci + Setrika (Wash + Iron)
  - Cuci Saja (Wash Only)
  - Setrika Saja (Iron Only)
- **Express Service** (express)
  - Cuci + Setrika (Wash + Iron)
  - Cuci Saja (Wash Only)
  - Setrika Saja (Iron Only)

**Items Covered:**
- Kemeja (Shirt)
- Celana (Pants)
- Jaket (Jacket)
- Selimut (Blanket)
- Sprei (Bed Sheet)

**Seeder Features:**
- ✅ Prevents duplicate entries
- ✅ Automatically checks existing data
- ✅ Safe to run multiple times
- ✅ Provides detailed logging

#### 5. Frontend Setup

```bash
# Navigate to frontend directory
cd ../frontend/frontend

# Install dependencies
npm install

# Configure API endpoint if needed
# Update API_BASE_URL in your JavaScript files
```

### Running the Application

#### Development Mode

**Terminal 1 - Backend:**
```bash
cd backend
go run cmd/main.go
```
Backend will run on `http://localhost:8080`

**Terminal 2 - Frontend:**
```bash
cd frontend/frontend
npm run dev
```
Frontend will run on `http://localhost:5173` (or next available port)

#### Production Build

**Backend:**
```bash
cd backend
go build -o chronos-laundry cmd/main.go
./chronos-laundry
```

**Frontend:**
```bash
cd frontend/frontend
npm run build
npm run preview
```

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/api/auth/login` | Admin login | No |
| POST | `/api/auth/register` | Admin registration | No |

### Transaction Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/transactions` | Get all transactions | Yes |
| GET | `/api/transactions/:id` | Get transaction by ID | Yes |
| POST | `/api/transactions` | Create new transaction | Yes |
| PUT | `/api/transactions/:id` | Update transaction | Yes |
| DELETE | `/api/transactions/:id` | Delete transaction | Yes |
| GET | `/api/transactions/track/:code` | Track by transaction code | No |

### Service Price Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/service-prices` | Get all service prices | Yes |
| GET | `/api/service-prices/:id` | Get service price by ID | Yes |
| POST | `/api/service-prices` | Create service price | Yes |
| PUT | `/api/service-prices/:id` | Update service price | Yes |
| DELETE | `/api/service-prices/:id` | Delete service price | Yes |

### Request/Response Examples

#### Login Request
```json
POST /api/auth/login
{
  "username": "admin",
  "password": "password123"
}
```

#### Login Response
```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "admin": {
      "id": 1,
      "username": "admin",
      "name": "Administrator"
    }
  }
}
```

#### Create Transaction Request
```json
POST /api/transactions
Authorization: Bearer <token>
{
  "customer_name": "John Doe",
  "customer_phone": "081234567890",
  "items": [
    {
      "service_price_id": 1,
      "quantity": 5,
      "notes": "Extra care for delicate items"
    }
  ]
}
```

## 🔧 Configuration

### Environment Variables

Create a `.env` file in the `backend` directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=chronos_laundry

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_change_this_in_production

# Server Configuration
PORT=8080
GIN_MODE=release  # Use 'debug' for development
```

### Frontend Configuration

Update API endpoint in JavaScript files if needed:

```javascript
// Example: js/config.js (create if needed)
const API_BASE_URL = 'http://localhost:8080/api';
```

## 🧪 Testing

```bash
# Run backend tests
cd backend
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

### How to Contribute

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/AmazingFeature
   ```
3. **Commit your changes**
   ```bash
   git commit -m 'Add some AmazingFeature'
   ```
4. **Push to the branch**
   ```bash
   git push origin feature/AmazingFeature
   ```
5. **Open a Pull Request**

### Coding Standards

- Follow Go best practices and conventions
- Use meaningful variable and function names
- Write clear commit messages
- Add comments for complex logic
- Ensure all tests pass before submitting PR

### Code Review Process

1. All PRs require at least one review
2. CI/CD checks must pass
3. Code must follow project style guidelines
4. Documentation must be updated if needed

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```
MIT License

Copyright (c) 2025 Chronos Laundry Team

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
```

## 👥 Team

- **Contributors**: [View all contributors](https://github.com/RidwanRamdhani/chronos-laundry/graphs/contributors)

## 💬 Support & Contact

- 📧 **Email**: ridwanramdhani@student.telkomuniversity.ac.id, daffarkananta10@gmail.com ,sanubarilegawa@student.telkomuniversity.ac.id
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/RidwanRamdhani/chronos-laundry/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/RidwanRamdhani/chronos-laundry/discussions)
- 📖 **Documentation**: [Wiki](https://github.com/RidwanRamdhani/chronos-laundry/wiki)

## 🗺️ Roadmap

### Version 1.0 (Current)
- ✅ Core transaction management
- ✅ Service price management
- ✅ JWT authentication
- ✅ Real-time tracking

### Version 1.1 (Planned)
- 🔄 Payment integration
- 🔄 SMS/Email notifications
- 🔄 Advanced reporting
- 🔄 Multi-branch support

### Version 2.0 (Future)
- 📱 Mobile application
- 🤖 AI-powered scheduling
- 📊 Advanced analytics dashboard
- 🌐 Multi-language support

## 📊 Project Status

![GitHub last commit](https://img.shields.io/github/last-commit/RidwanRamdhani/chronos-laundry)
![GitHub issues](https://img.shields.io/github/issues/RidwanRamdhani/chronos-laundry)
![GitHub pull requests](https://img.shields.io/github/issues-pr/RidwanRamdhani/chronos-laundry)

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library for Go
- [Vite](https://vitejs.dev/) - Next generation frontend tooling
- [JWT](https://jwt.io/) - JSON Web Tokens



<div align="center">

**[Website](https://chronoslaundry.com)** • **[Documentation](https://docs.chronoslaundry.com)** • **[Demo](https://demo.chronoslaundry.com)**

Made with ❤️ by the Chronos Laundry Team

⭐ Star us on GitHub — it motivates us a lot!

</div>
