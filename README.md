# 🏦 Banking System API

<div align="center">

![Go](https://img.shields.io/badge/Go-1.24.5-blue?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-green?style=for-the-badge&logo=gin)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-blue?style=for-the-badge&logo=postgresql)
![JWT](https://img.shields.io/badge/JWT-Authentication-orange?style=for-the-badge&logo=jsonwebtokens)
![GORM](https://img.shields.io/badge/GORM-ORM-red?style=for-the-badge)

*A modern, secure, and scalable banking system built with Go, featuring JWT authentication, account management, and money transfer capabilities.*

</div>

## ✨ Features

### 🔐 **Authentication & Security**
- **JWT-based Authentication** with access and refresh tokens
- **Password Hashing** using bcrypt for secure password storage
- **Token Refresh** mechanism for seamless user experience
- **Middleware Protection** for secure API endpoints
- **Session Management** with secure cookie handling

### 💰 **Banking Operations**
- **Account Creation** and management
- **Money Transfers** between accounts with transaction safety
- **Balance Checking** and account information retrieval
- **Database Transactions** ensuring data consistency
- **Insufficient Funds Protection** with proper error handling

### 🛡️ **Data Safety**
- **Database Locking** during transfers to prevent race conditions
- **Atomic Transactions** ensuring all-or-nothing operations
- **Input Validation** and error handling
- **PostgreSQL** for robust data storage

## 🚀 Quick Start

### Prerequisites
- Go 1.24.5 or higher
- PostgreSQL database
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/BankingSystem.git
   cd BankingSystem
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Environment Setup**
   Create a `.env` file in the root directory:
   ```env
   DB_url=postgres://username:password@localhost:5432/banking_db
   SECRET=your-super-secret-jwt-key
   PORT=8080
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080` (or your specified PORT)

## 📚 API Documentation

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/signup` | Register a new user | ❌ |
| `POST` | `/login` | User login | ❌ |
| `POST` | `/refreshToken` | Refresh access token | ❌ |
| `GET` | `/logout` | User logout | ❌ |

### Banking Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/api/post` | Create account (from) | ✅ |
| `POST` | `/api/postt` | Create account (to) | ✅ |
| `POST` | `/api/transfer` | Transfer money | ✅ |
| `GET` | `/api/get/:id` | Get account details | ✅ |

### Request Examples

#### 🔐 User Registration
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword123"
  }'
```

#### 🔑 User Login
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword123"
  }'
```

#### 💰 Create Account
```bash
curl -X POST http://localhost:8080/api/post \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -d '{
    "user_id": 1,
    "owner": "John Doe",
    "balance": 1000.00
  }'
```

#### 💸 Transfer Money
```bash
curl -X POST http://localhost:8080/api/transfer \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN" \
  -d '{
    "from_id": 1,
    "to_id": 2,
    "amount": 250.50
  }'
```

## 🏗️ Project Structure

```
BankingSystem/
├── Controller/           # API controllers
│   ├── autho.go         # Authentication logic
│   ├── controller.go    # Banking operations
│   └── refreshcontro.go # Token refresh
├── intializers/         # App initialization
│   ├── database.go      # DB connection
│   ├── LoadEnvVar.go    # Environment variables
│   └── Syncdatabase.go  # DB migration
├── middleware/          # Middleware functions
│   └── Reqautho.go      # JWT authentication
├── models/              # Data models
│   ├── model.go         # Account model
│   ├── refresh.go       # Refresh token model
│   └── user.go          # User model
├── utils/               # Utility functions
│   └── jwt.go           # JWT operations
├── main.go              # Application entry point
├── go.mod               # Go modules
└── go.sum               # Dependencies
```

## 🔧 Technologies Used

- **Backend**: Go 1.24.5
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (golang-jwt/jwt)
- **Password Hashing**: bcrypt
- **Environment**: godotenv

## 🛡️ Security Features

- **JWT Tokens**: Secure authentication with configurable expiration
- **Password Hashing**: bcrypt with salt rounds
- **Database Transactions**: ACID compliance for money transfers
- **Row-level Locking**: Prevents concurrent modification issues
- **Input Validation**: Comprehensive request validation
- **Error Handling**: Secure error responses without sensitive data

## 🚦 Error Handling

The API provides comprehensive error handling:

- **400 Bad Request**: Invalid input data
- **401 Unauthorized**: Missing or invalid authentication
- **404 Not Found**: Account not found
- **500 Internal Server Error**: Server-side errors

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Divyansh Tiwari**
- GitHub: [@yourusername](https://github.com/yourusername)
- Email: your.email@example.com

---

<div align="center">

**⭐ Star this repository if you found it helpful!**

Made with ❤️ and Go

</div>
