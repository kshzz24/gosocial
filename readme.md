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
- **Email:** SMTP (Gmail/SendGrid)
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
│   │   ├── auth.go              # Authentication handlers
│   │   └── subreddit.go         # Subreddit handlers
│   ├── middleware/
│   │   └── auth.go              # JWT authentication middleware
│   ├── models/
│   │   ├── user.go              # User model and database operations
│   │   └── subreddit.go         # Subreddit model and database operations
│   └── utils/
│       ├── jwt.go               # JWT generation and validation
│       ├── password.go          # Password hashing utilities
│       ├── token.go             # Reset token generation
│       └── email.go             # Email sending utilities
├── migrations/                   # SQL migration files
│   ├── 001_create_users_table.sql
│   ├── 002_add_reset_token_to_users.sql
│   └── 003_create_subreddits_table.sql
├── .env                         # Environment variables (not in git)
├── .gitignore
├── go.mod
└── go.sum
```

## ✅ Completed Features (Phase 1: Complete Authentication System)

### 1. Project Setup

- [x] Go module initialization
- [x] Project structure with clean architecture
- [x] PostgreSQL database setup with pgAdmin
- [x] Environment variable configuration
- [x] Dependency management (Gin, JWT, bcrypt, gomail)

### 2. Database Layer

- [x] PostgreSQL connection module
- [x] Users table with complete schema
- [x] Password reset token fields
- [x] Database connection pooling
- [x] Graceful connection closing
- [x] Database indexes for performance

### 3. Core Utilities

- [x] Password hashing with bcrypt
- [x] Password verification
- [x] JWT token generation (HS256)
- [x] JWT token validation
- [x] Token claims extraction
- [x] Secure reset token generation
- [x] Email sending via SMTP

### 4. User Model

- [x] User struct with all fields
- [x] CreateUser function
- [x] GetUserByEmail function
- [x] GetUserByID function
- [x] GetUserByResetToken function
- [x] UpdatePassword function
- [x] SaveResetToken function
- [x] ClearResetToken function
- [x] Proper error handling

### 5. Authentication Handlers (Complete)

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

- [x] **POST /api/logout** - User logout

  - Authentication required
  - Clean logout flow

- [x] **PUT /api/change-password** - Change password

  - Authentication required
  - Old password verification
  - New password validation
  - Password update

- [x] **POST /auth/forgot-password** - Request password reset

  - Email validation
  - Reset token generation (1-hour expiry)
  - Email delivery with reset link
  - Security: same response for valid/invalid emails

- [x] **POST /auth/reset-password** - Reset password with token
  - Token validation
  - Expiry checking
  - Password update
  - Token invalidation after use

### 6. Authentication Middleware

- [x] RequireAuth() - Enforces JWT authentication

  - Token extraction from Authorization header
  - Token validation
  - User context injection
  - 401 responses for invalid/missing tokens

- [x] OptionalAuth() - Optional authentication
  - Allows anonymous access
  - Injects user context when token present

### 7. Email System

- [x] SMTP configuration (Gmail/SendGrid support)
- [x] HTML email templates
- [x] Password reset email with secure links
- [x] Professional email formatting

### 8. Server Configuration

- [x] Gin router setup
- [x] Route grouping (public vs protected)
- [x] Middleware integration
- [x] Logging configuration
- [x] Configurable port
- [x] Error handling patterns

## ✅ Phase 2A: Subreddits System (COMPLETED)

### 1. Database Schema

- [x] Subreddits table with advanced features
- [x] JSONB support for rules and flairs
- [x] Foreign key relationships to users
- [x] Indexes for performance (name, created_by, members_count)
- [x] Check constraints for name validation
- [x] Support for NSFW and private communities

### 2. Subreddit Model

- [x] Subreddit struct with all fields including JSONB
- [x] CreateSubreddit function with JSONB handling
- [x] GetSubredditByName function
- [x] GetSubredditByID function
- [x] ListSubreddits with pagination
- [x] UpdateSubreddit function
- [x] Proper error handling and validation

### 3. Subreddit Handlers (Complete CRUD)

- [x] **POST /api/subreddits** - Create subreddit (auth required)
  - User authentication validation
  - Name format validation (lowercase, numbers, underscores only)
  - Duplicate name detection
  - Creator automatically becomes first member
  - Default JSONB arrays for rules and flairs
- [x] **GET /api/subreddits/:name** - Get subreddit by name (public)
  - URL parameter extraction
  - 404 handling for non-existent subreddits
  - Complete subreddit data return
- [x] **GET /api/subreddits** - List all subreddits (public)
  - Dual pagination support (limit/offset AND page/per_page)
  - Ordered by popularity (members_count DESC)
  - Pagination metadata in response
  - Max limit enforcement (100)
- [x] **PUT /api/subreddits/:id** - Update subreddit (auth required, owner only)
  - Ownership verification
  - Update display name, description, rules, images, flags
  - JSONB field updates
  - Protected fields (name, created_by, counts)
- [x] **DELETE /api/subreddits/:id** - Delete subreddit (auth required, owner only)
  - Ownership verification
  - Cascade deletion handling

### 4. Advanced Features Implemented

- [x] JSONB support for dynamic rules and flairs
- [x] Name validation with regex patterns
- [x] Support for NSFW and private subreddits
- [x] Image URL fields (banner and icon) - ready for future uploads
- [x] Flexible pagination (offset-based and page-based)
- [x] Members count tracking
- [x] Active users field (prepared for real-time features)

## 🚀 API Endpoints (Current)

### Authentication (Public Routes)

#### Register New User

```http
POST /auth/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "SecurePass123!"
}

Response 200:
{
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "created_at": "2025-11-06T..."
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Login

```http
POST /auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "SecurePass123!"
}

Response 200:
{
  "user": { ... },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Forgot Password

```http
POST /auth/forgot-password
Content-Type: application/json

{
  "email": "john@example.com"
}

Response 200:
{
  "message": "If that email exists, a reset link has been sent"
}

Note: Reset link sent to email with 1-hour expiry
```

#### Reset Password

```http
POST /auth/reset-password
Content-Type: application/json

{
  "token": "reset_token_from_email",
  "new_password": "NewSecurePass123!"
}

Response 200:
{
  "message": "Password reset successfully"
}
```

### User Routes (Protected - Require JWT Token)

#### Get Current User

```http
GET /api/me
Authorization: Bearer <jwt_token>

Response 200:
{
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "created_at": "2025-11-06T..."
  }
}
```

#### Logout

```http
POST /api/logout
Authorization: Bearer <jwt_token>

Response 200:
{
  "message": "Logged out successfully"
}
```

#### Change Password

```http
PUT /api/change-password
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "old_password": "SecurePass123!",
  "new_password": "NewSecurePass456!"
}

Response 200:
{
  "message": "Password updated successfully"
}
```

## Subreddit Routes

### Create Subreddit (Protected)

```http
POST /api/subreddits
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "name": "golang",
  "display_name": "Golang",
  "description": "The Go programming language community",
  "rules": [
    {
      "title": "Be Respectful",
      "description": "No toxicity or harassment"
    }
  ],
  "is_nsfw": false,
  "is_private": false
}

Response 201:
{
  "message": "Subreddit created successfully",
  "subreddit": {
    "id": 1,
    "name": "golang",
    "display_name": "Golang",
    "description": "The Go programming language community",
    "rules": [...],
    "created_by": 5,
    "members_count": 1,
    "is_nsfw": false,
    "is_private": false,
    "created_at": "2025-11-06T...",
    "updated_at": "2025-11-06T..."
  }
}
```

### Get Subreddit (Public)

```http
GET /api/subreddits/:name

Example: GET /api/subreddits/golang

Response 200:
{
  "subreddit": {
    "id": 1,
    "name": "golang",
    "display_name": "Golang",
    "description": "...",
    "rules": [...],
    "members_count": 150,
    "active_users": 12,
    "created_at": "2025-11-06T..."
  }
}

Response 404:
{
  "error": "Subreddit not found"
}
```

### List Subreddits (Public)

```http
GET /api/subreddits?limit=20&offset=0
# OR
GET /api/subreddits?page=2&per_page=15

Response 200:
{
  "subreddits": [
    {
      "id": 1,
      "name": "golang",
      "display_name": "Golang",
      "members_count": 150,
      ...
    },
    {
      "id": 2,
      "name": "programming",
      "display_name": "Programming",
      "members_count": 1250,
      ...
    }
  ],
  "pagination": {
    "limit": 20,
    "offset": 0,
    "count": 2
  }
}
```

### Update Subreddit (Protected)

```http
PUT /api/subreddits/:id
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "display_name": "Golang Programming",
  "description": "Updated description",
  "rules": [
    {
      "title": "New Rule",
      "description": "New rule description"
    }
  ],
  "is_nsfw": false
}

Response 200:
{
  "message": "Subreddit updated successfully",
  "subreddit": { ... }
}

Response 403:
{
  "error": "You can only update subreddits you created"
}
```

### Delete Subreddit (Protected)

```http
DELETE /api/subreddits/:id
Authorization: Bearer <jwt_token>

Response 200:
{
  "message": "Subreddit deleted successfully"
}

Response 403:
{
  "error": "You can only delete subreddits you created"
}
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

# Email Configuration (SMTP)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-16-char-app-password
SMTP_FROM=your-email@gmail.com

# Frontend Configuration
FRONTEND_URL=http://localhost:3000

# Server Configuration
PORT=8080
```

**Note for Gmail:** You need to generate an App Password:

1. Go to Google Account → Security → 2-Step Verification
2. Search for "App passwords"
3. Generate password for "Mail"
4. Use the 16-character password in SMTP_PASSWORD

5. **Run migrations**

```bash
# Migration 1: Create users table
psql -U your_db_user -d gosocial -f migrations/001_create_users_table.sql

# Migration 2: Add password reset functionality
psql -U your_db_user -d gosocial -f migrations/002_add_reset_token_to_users.sql

# Migration 3: Create subreddits table with JSONB support
psql -U your_db_user -d gosocial -f migrations/003_create_subreddits_table.sql
```

6. **Run the server**

```bash
go run cmd/api/main.go
```

Server will start on `http://localhost:8080` 🚀

## 📋 Roadmap

### ✅ Phase 1: Complete Authentication System (COMPLETED)

- [x] Database setup with PostgreSQL
- [x] User registration with validation
- [x] User login with JWT tokens
- [x] Protected routes with middleware
- [x] Get current user endpoint
- [x] Logout functionality
- [x] Change password (authenticated)
- [x] Forgot password with email
- [x] Reset password with token
- [x] Email system integration
- [x] Security best practices

### ✅ Phase 2A: Subreddits System (COMPLETED)

- [x] Subreddits table with JSONB support
- [x] Subreddit model with all CRUD operations
- [x] Create subreddit (auth required, owner only)
- [x] Get subreddit by name (public)
- [x] List subreddits with pagination (public)
- [x] Update subreddit (auth required, owner only)
- [x] Delete subreddit (auth required, owner only)
- [x] Name validation and format checking
- [x] Support for NSFW and private communities
- [x] JSONB support for rules and flairs
- [x] Dual pagination support (offset and page-based)

### 🔄 Phase 2B: Posts System (IN PROGRESS)

- [ ] Posts table and model
- [ ] Create post endpoint
- [ ] List posts with pagination
- [ ] Get single post
- [ ] Update post (author only)
- [ ] Delete post (author only)
- [ ] Post voting system

### 📅 Phase 3: Comments

- [ ] Comments table and model
- [ ] Add comment to post
- [ ] Nested comments structure
- [ ] Comment voting
- [ ] Edit/delete comments

### 📅 Phase 4: User Profiles & Karma

- [ ] User profile endpoints
- [ ] User karma calculation
- [ ] User post/comment history
- [ ] Follow/unfollow users

### 📅 Phase 5: Subreddit Membership & Moderation

- [ ] Join/leave subreddit tracking
- [ ] Subreddit moderators system
- [ ] Moderation tools and permissions
- [ ] User banning and content removal

### 📅 Phase 6: Advanced Features

- [ ] Real-time notifications (WebSockets)
- [ ] Image upload for posts and subreddits
- [ ] Search functionality with Elasticsearch
- [ ] Content reporting system
- [ ] Awards and gilding

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
- ✅ Email integration with SMTP
- ✅ Token-based password reset flow
- ✅ Security best practices (timing attacks, email enumeration)
- ✅ JSONB in PostgreSQL for flexible data structures
- ✅ Complex database queries with foreign keys
- ✅ Pagination patterns (offset-based and page-based)
- ✅ Input validation and regex patterns
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

**Kushagra** - [GitHub](https://github.com/kshzz24)

## 🙏 Acknowledgments

- Go documentation and community
- Gin framework documentation
- PostgreSQL documentation
- Various Go learning resources and tutorials

---

**Current Status:** Phase 1 ✅ | Phase 2A ✅ | Phase 2B In Progress 🚀

**Last Updated:** November 7, 2025

**Next Milestone:** Posts CRUD operations with voting system

---

## 📊 Project Statistics

- **Total Endpoints:** 12 (7 auth + 5 subreddit)
  - 4 public auth routes
  - 3 protected auth routes
  - 3 public subreddit routes
  - 2 protected subreddit routes
- **Database Tables:** 2 (users, subreddits)
- **Authentication Methods:** JWT with email-based password reset
- **Security Features:** bcrypt hashing, token expiry, input validation, ownership checks
- **Advanced Features:** JSONB support, dual pagination, regex validation

## 🎯 Key Technical Achievements

- **Complete Authentication Flow** - Registration, login, password management
- **Email Integration** - SMTP with HTML templates for password reset
- **Advanced Database Design** - Foreign keys, JSONB, indexes, constraints
- **Flexible Pagination** - Support for both offset-based and page-based
- **Input Validation** - Regex patterns, length checks, format validation
- **Ownership Authorization** - Creator-only updates and deletions
- **Clean Architecture** - Separation of concerns (handlers, models, middleware, utils)
