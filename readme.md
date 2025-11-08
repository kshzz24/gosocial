# GoSocial - Reddit Clone Backend

A Reddit-style social platform backend built with Go, PostgreSQL, and JWT authentication.

## 🚀 Tech Stack

- **Go 1.24** | **Gin** | **PostgreSQL** | **JWT** | **bcrypt** | **SMTP**

## 📁 Project Structure

```
gosocial/
├── cmd/api/main.go
├── internal/
│   ├── database/postgres.go
│   ├── handlers/
│   │   ├── auth.go
│   │   └── subreddit.go
│   ├── middleware/auth.go
│   ├── models/
│   │   ├── user.go
│   │   └── subreddit.go
│   └── utils/
│       ├── jwt.go
│       ├── password.go
│       ├── token.go
│       └── email.go
└── migrations/
    ├── 001_create_users_table.sql
    ├── 002_add_reset_token_to_users.sql
    └── 003_create_subreddits_table.sql
```

## ✅ Completed Features

### Authentication System
- User registration & login with JWT
- Password management (change, forgot, reset via email)
- Protected routes with middleware
- Email integration (SMTP)

### Subreddits System  
- Full CRUD operations
- JSONB support for rules/flairs
- Dual pagination (offset & page-based)
- NSFW & private community support
- Ownership-based authorization

## 🔧 Quick Start

### 1. Setup Database
```bash
createdb gosocial
psql -d gosocial -f migrations/001_create_users_table.sql
psql -d gosocial -f migrations/002_add_reset_token_to_users.sql
psql -d gosocial -f migrations/003_create_subreddits_table.sql
```

### 2. Configure Environment
Create `.env`:
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=gosocial
DB_SSLMODE=disable

# JWT
JWT_SECRET=your_secret_key

# Email (Gmail)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_FROM=your-email@gmail.com
FRONTEND_URL=http://localhost:3000

# Server
PORT=8080
```

### 3. Run
```bash
go mod download
go run cmd/api/main.go
```

Server starts at `http://localhost:8080` 🚀

## 📡 API Endpoints

### Auth (Public)
| Method |        Endpoint        | Description |
|--------|----------|-------------|
| POST |      `/auth/register`    | Create account |
| POST |      `/auth/login`       | Login user     |
| POST | `/auth/forgot-password`  | Request reset |
| POST | `/auth/reset-password`   | Reset password |

### Auth (Protected)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/me` | Get current user |
| POST | `/api/logout` | Logout |
| PUT | `/api/change-password` | Change password |

### Subreddits
| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/api/subreddits` | ✅ | Create subreddit |
| GET | `/api/subreddits` | ❌ | List all (paginated) |
| GET | `/api/subreddits/:name` | ❌ | Get by name |
| PUT | `/api/subreddits/:id` | ✅ | Update (owner only) |
| DELETE | `/api/subreddits/:id` | ✅ | Delete (owner only) |

## 📝 Example Requests

### Create Subreddit
```bash
POST /api/subreddits
Authorization: Bearer <token>

{
  "name": "golang",
  "display_name": "Golang",
  "description": "Go programming community",
  "is_nsfw": false,
  "is_private": false
}
```

### List Subreddits (Pagination)
```bash
# Offset-based
GET /api/subreddits?limit=20&offset=0

# Page-based
GET /api/subreddits?page=2&per_page=15
```

## 📋 Roadmap

- ✅ **Phase 1:** Authentication System
- ✅ **Phase 2A:** Subreddits CRUD
- 🔄 **Phase 2B:** Posts System (In Progress)
- 📅 **Phase 3:** Comments & Nested Replies
- 📅 **Phase 4:** Voting System
- 📅 **Phase 5:** User Profiles & Karma
- 📅 **Phase 6:** Image Uploads & Search
- 📅 **Phase 7:** Real-time Features & Optimization

## 📊 Current Stats

- **Endpoints:** 12 (7 auth + 5 subreddits)
- **Tables:** 2 (users, subreddits)
- **Features:** JWT auth, email reset, JSONB, pagination, ownership checks

## 👤 Author

**Kshitiz** - [GitHub](https://github.com/kshzz24)

---

**Status:** Phase 2A Complete ✅ | Building Posts Next 🚀