# GoSocial - Reddit Clone Backend

A Reddit-style social media platform backend built with Go, focusing on learning Go best practices, clean architecture, and RESTful API design.

## 🎯 Project Overview

GoSocial is a learning project to master Go programming by building a feature-rich social platform from the ground up. The project follows a 16-week development roadmap progressing from basic authentication to advanced features like real-time updates and content moderation.

## 🏗️ Tech Stack

- **Language:** Go 1.24
- **Web Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT (HS256)
- **Password Hashing:** bcrypt
- **Environment Management:** godotenv

## 📁 Project Structure

```
gosocial/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── database/
│   │   └── postgres.go          # Database connection logic
│   ├── handlers/
│   │   └── auth.go              # Authentication handlers
│   ├── middleware/
│   │   └── auth.go              # JWT authentication middleware
│   ├── models/
│   │   └── user.go              # User model and database operations
│   └── utils/
│       ├── jwt.go               # JWT generation and validation
│       └── password.go          # Password hashing utilities
├── migrations/                   # SQL migration files
├── .env                         # Environment variables (not in git)
├── .gitignore
├── go.mod
└── go.sum
```

## ✅ Completed Features (Phase 1: Authentication)

### 1. Project Setup

- [x] Go module initialization
- [x] Project structure with clean architecture
- [x] PostgreSQL database setup with pgAdmin
- [x] Environment variable configuration
- [x] Dependency management

### 2. Database Layer

- [x] PostgreSQL connection module
- [x] Users table with proper schema
- [x] Database connection pooling
- [x] Graceful connection closing

### 3. Core Utilities

- [x] Password hashing with bcrypt
- [x] Password verification
- [x] JWT token generation (HS256)
- [x] JWT token validation
- [x] Token claims extraction

### 4. User Model

- [x] User struct with proper types
- [x] CreateUser function
- [x] GetUserByEmail function
- [x] GetUserByID function
- [x] Proper error handling

### 5. Authentication Handlers

- [x] **POST /auth/register** - User registration
  - Input validation
  - Duplicate email check
  - Password hashing
  - JWT token generation
  - Secure response (no password in output)
- [x] **POST /auth/login** - User login
  - Email/password validation
  - Password verification
  - JWT token generation
  - Generic error messages for security
- [x] **GET /api/me** - Get current user
  - JWT authentication required
  - User data retrieval
  - Secure response format

### 6. Authentication Middleware

- [x] RequireAuth() - Enforces JWT authentication

  - Token extraction from Authorization header
  - Token validation
  - User context injection
  - 401 responses for invalid/missing tokens

- [x] OptionalAuth() - Optional authentication
  - Allows anonymous access
  - Injects user context when token present

### 7. Server Configuration

- [x] Gin router setup
- [x] Route grouping (public vs protected)
- [x] Middleware integration
- [x] Logging configuration
- [x] Configurable port

## 🚀 API Endpoints (Current)

### Authentication (Public)

```http
POST /auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "SecurePass123!"
}

Response: { "user": {...}, "token": "jwt_token" }
```

```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePass123!"
}

Response: { "user": {...}, "token": "jwt_token" }
```

### User (Protected)

```http
GET /api/me
Authorization: Bearer <jwt_token>

Response: { "user": {...} }
```

## 🔧 Setup & Installation

### Prerequisites

- Go 1.24 or higher
- PostgreSQL 12+
- pgAdmin (optional, for database management)

### Installation Steps

1. **Clone the repository**

```bash
git clone https://github.com/kshzz24/gosocial.git
cd gosocial
```

2. **Install dependencies**

```bash
go mod download
```

3. **Setup PostgreSQL database**

```sql
CREATE DATABASE gosocial;
```

4. **Create `.env` file**

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=gosocial
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_change_this_in_production

# Server Configuration
PORT=8080
```

5. **Run migrations**

```bash
# Create users table
psql -U your_db_user -d gosocial -f migrations/001_create_users_table.sql
```

6. **Run the server**

```bash
go run cmd/api/main.go
```

Server will start on `http://localhost:8080` 🚀

## 📋 Roadmap

### ✅ Phase 1: Foundation & Authentication (COMPLETED)

- Database setup
- User registration
- User login
- JWT authentication
- Protected routes

### 🔄 Phase 2: Posts & Content (IN PROGRESS)

- [ ] Posts table and model
- [ ] Create post endpoint
- [ ] List posts with pagination
- [ ] Get single post
- [ ] Update post (author only)
- [ ] Delete post (author only)
- [ ] Post voting system

### 📅 Phase 3: Subreddits

- [ ] Subreddits table and model
- [ ] Create subreddit
- [ ] Join/leave subreddit
- [ ] List subreddit posts
- [ ] Subreddit moderators

### 📅 Phase 4: Comments

- [ ] Comments table and model
- [ ] Add comment to post
- [ ] Nested comments structure
- [ ] Comment voting
- [ ] Edit/delete comments

### 📅 Phase 5: User Profiles & Karma

- [ ] User profile endpoints
- [ ] User karma calculation
- [ ] User post/comment history
- [ ] Follow/unfollow users

### 📅 Phase 6: Advanced Features

- [ ] Real-time notifications (WebSockets)
- [ ] Image upload for posts
- [ ] Search functionality
- [ ] Moderation tools
- [ ] Report system

### 📅 Phase 7: Performance & Scale

- [ ] Redis caching
- [ ] Rate limiting
- [ ] Database indexing optimization
- [ ] Goroutines for async tasks
- [ ] Background job processing

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/handlers
```

## 📚 Learning Goals

This project focuses on mastering:

- ✅ Go project structure and organization
- ✅ RESTful API design
- ✅ Database operations with PostgreSQL
- ✅ JWT authentication implementation
- ✅ Middleware patterns
- ✅ Error handling in Go
- ✅ Password security with bcrypt
- 🔄 Concurrent programming with goroutines
- 🔄 Channels for communication
- 🔄 Context for cancellation
- 🔄 Performance optimization
- 🔄 Testing best practices

## 🤝 Contributing

This is a learning project, but suggestions and improvements are welcome!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is for educational purposes.

## 👤 Author

**Kshitiz Bartaria** - [GitHub](https://github.com/kshzz24)

## 🙏 Acknowledgments

- Go documentation and community
- Gin framework documentation
- PostgreSQL documentation
- Various Go learning resources and tutorials

---

**Current Status:** Phase 1 Complete ✅ | Phase 2 Starting 🚀

**Next Milestone:** Posts CRUD operations with voting system
