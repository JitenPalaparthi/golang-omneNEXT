# Fiber + GORM + Zerolog + Postgres — Users CRUD

Minimal Go backend using:
- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/) with Postgres
- [Zerolog](https://github.com/rs/zerolog)
- [godotenv](https://github.com/joho/godotenv)

## Features
- Users CRUD: `id, name, email, password (hashed), status, last_modified (unix)`
- Pagination on GET /users: `limit` & `skip`
- Auto-migrate on startup
- CORS enabled for local Vite (`http://localhost:5173`)
- Sensible error handling + JSON responses

## Quick start

### 1) Start Postgres (Docker)
```bash
docker compose up -d
# DB on localhost:5432 user=postgres password=postgres db=usersdb
```

### 2) Configure env and run
```bash
cp .env.example .env
# Optionally edit DATABASE_URL / CORS_ORIGIN
go run ./cmd/server
```

### 3) Endpoints
- `POST   /users`
- `GET    /users?limit=10&skip=0`
- `GET    /users/:id`
- `PUT    /users/:id`
- `DELETE /users/:id`

`last_modified` is updated automatically on create/update.

### 4) cURL quick tests
See [`scripts/curl.sh`](scripts/curl.sh) or copy below.

## Build
```bash
go build -o bin/server ./cmd/server
./bin/server
```

## Notes
- Default port: **58080** (`PORT` env).
- Password is **hashed** with bcrypt (`golang.org/x/crypto/bcrypt`).
- Email is `UNIQUE` in DB.
