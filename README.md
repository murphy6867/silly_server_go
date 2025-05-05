# Chirps Server Go

A simple RESTful microservice written in Go, designed as an example project to showcase my backend engineering skills for job applications. It demonstrates clean project structure, database integration, API design, configuration management, and static asset serving.

---

## Features

- **Health & Metrics**

  - `GET /api/healthz` – simple health check
  - `GET /api/metrics` – Prometheus-style metrics

- **Database Reset**

  - `POST /api/reset` – reset the database to a clean state (for demos and testing)

- **User Management**

  - `POST /api/users` – create a new user

- **Chirp Management**

  - `GET  /api/chirps` – list all chirps
  - `GET  /api/chirps/{id}` – fetch a single chirp by ID
  - `POST /api/chirps` – create a new chirp

- **Admin Endpoints** (protected by a separate path)

  - `GET  /admin/metrics`
  - `POST /admin/reset`

- **Static File Serving**
  - Hosts a simple front-end under `/app/*`
  - Serves images under `/api/assets/`

---

## Tech Stack

- **Language & Framework**

  - Go’s standard `net/http` for routing & handlers
  - Modular “clean architecture” layout:
    - `internal/handler` – HTTP handlers & middleware
    - `internal/database` – SQLC-generated, type-safe queries
    - `internal/user` & `internal/chirp` – domain services & repositories

- **Data Storage**

  - PostgreSQL via `database/sql` + `github.com/lib/pq`
  - SQLC for compile-time checked queries (`sqlc.yml` + `sql/`)

- **Configuration**

  - Environment variables with `github.com/joho/godotenv`
  - `.env.example` for local setup

- **Observability**

  - Simple metrics middleware to count requests
  - Readiness & liveness endpoints

- **Static Assets**
  - `web/static/` folder served via `http.FileServer`

---

## Getting Started

1. **Clone the repo**

   ```bash
   git clone https://github.com/murphy6867/silly_server_go.git
   cd silly_server_go

   ```

2. \*\*Configure environment
   cp .env.example .env

   # Edit .env → set DB_URL=postgres://user:pass@localhost:5432/sillydb?sslmode=disable

3. \*\*Prepare the database
   Create a PostgreSQL database named sillydb (or whatever you set in DB_URL)
   Run the SQL schema/migrations in the sql/ folder:
   `psql $DB_URL < sql/schema.sql`

4. Generate Go DB code
   `sqlc generate`

5. Build & Run
   `go build -o silly-server ./cmd/chirps`
   `./silly-server`
   # Server listens on :8080

--

## Usage Examples

- Health check
  `curl http://localhost:8080/api/healthz`
- Create a user
  `curl -X POST http://localhost:8080/api/users \`
  `-H "Content-Type: application/json" \`
  `-d '{"username":"alice","email":"alice@example.com"}'`
- Post a chirp
  `curl -X POST http://localhost:8080/api/chirps \`
  `-H "Content-Type: application/json" \`
  `-d '{"user_id":1,"message":"Hello, world!"}'`
- List chirps
  `curl http://localhost:8080/api/chirps`

--

## Project Structure

├── cmd/
│ └── chirps/ # entrypoint: main.go
├── internal/
│ ├── handler/ # HTTP handlers & middleware
│ ├── database/ # SQLC-generated DB queries
│ ├── user/ # user service & repository
│ └── chirp/ # chirp service & repository
├── sql/ # raw SQL schema & queries
├── web/
│ └── static/ # frontend assets
├── .env.example # env var template
├── sqlc.yml # SQLC config
├── go.mod & go.sum # module dependencies
└── README.md # this file

--

## Contributing

This repository is intended as a standalone sample. Feel free to clone, fork, or adapt for your own learning or demonstrations.
