Golang Blog Platform

Features

- User registration/login with JWT
- Authenticated post creation/update/delete
- View all posts or a single post
- Middleware auth checks
- PostgreSQL via GORM

Technologies

- Go (`net/http`, `gorm`)
- PostgreSQL
- JWT + bcrypt
- Render for deployment

Setup

```bash
go mod tidy
cp .env.example .env
go run main.go

