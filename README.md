# Chirps Server Go
Chirps Server Go is a RESTful microservice implemented in Go, showcasing robust backend engineering practices. It leverages clean architecture to separate concerns (controllers, services, models) and integrates a PostgreSQL database using SQLC for type-safe query generation. Key features include JWT-based authentication, Prometheus metrics, health/readiness endpoints, configuration via environment variables, and static asset serving for a minimal front-end UI. This project demonstrates building scalable, maintainable APIs with industry-standard tooling.

## Key Features
- **RESTful API Endpoints**: Create and retrieve chirps (short posts) with a clear URL structure (e.g., `/api/v1/chirps`).
- **User Management & Authentication**: Secure user signup/login with password hashing and JWT tokens to protect sensitive routes.
- **Clean Architecture**: Organized codebase separating routing, business logic, and storage layers for maintainability.
- **Database Integration**: PostgreSQL backend with schema managed via SQLC code generation for efficient, type-safe queries.
- **Metrics & Health Checks**: Prometheus-formatted metrics endpoint (`/metrics`) and a `/healthz` readiness probe for monitoring.
- **Configuration Management**: Environment-based configuration (e.g., DB credentials, server port) loaded from a `.env` file.
- **Static Asset Serving**: Built-in HTTP server hosts static files (HTML/JS/CSS) to serve a simple front-end UI.

## Tech Stack
- **Go**: Language (Go 1.21+).
- **HTTP Router**: chi/v5 or Go net/http (for request routing).
- **PostgreSQL**: Relational database system.
- **SQLC**: SQL code generator for Go (type-safe database queries).
- **Prometheus**: Monitoring metrics exposition.
- **JWT**: JSON Web Tokens for authentication.
- **Viper/envconfig**: Environment-based configuration management.
- **Other**: Standard libraries (net/http, database/sql), logging, hashing utilities, etc.

## Setup
1. Clone the repository and navigate in:
     ```
     git clone https://github.com/yourusername/chirps-server-go.git
     cd chirps-server-go
     ```
2. Copy and edit the environment file:
     ```
       cp .env.example .env
     ```
   - Open .env and set the database connection and JWT secret:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=yourpassword
    DB_NAME=chirpsdb
    JWT_SECRET=your_jwt_secret_key
    ```
3. Install dependencies:
     ```
     go mod download
     ```
4. Prepare the database:
   - Ensure PostgreSQL is running.
   - Create the database (matching `DB_NAME` in `.env`):
       ```
       createdb chirpsdb
       ```
   - Run any provided migrations or use the reset utility to initialize schema:
       ```
       go run reset.go
       ```
7. Run the server: The server will start on the port defined in .env (default 8080).
   ```
   go run main.go
   ```
## Usage Examples
- Create a new user:
  ```
  curl -X POST http://localhost:8080/api/v1/users \
     -H "Content-Type: application/json" \
     -d '{"username":"alice","email":"alice@example.com","password":"password123"}'
  ```
  A successful response returns the new user data in JSON.
- User login (obtain JWT token):
  ```
  curl -X POST http://localhost:8080/api/v1/login \
     -H "Content-Type: application/json" \
     -d '{"username":"alice","password":"password123"}'
  ```
  The response includes a JWT token for authenticated requests.
- Create a chirp (authenticated):
  ```
  curl -X POST http://localhost:8080/api/v1/chirps \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <YOUR_JWT_TOKEN>" \
     -d '{"message":"Hello, world!"}'
  ```
  Replace <YOUR_JWT_TOKEN> with the token from the login step. Returns the created chirp object.
- Retrieve recent chirps:
  ```
  curl http://localhost:8080/api/v1/chirps
  ```
  Returns a JSON list of chirps with metadata (user, timestamp, etc.).
- Health check:
  ```
  curl http://localhost:8080/healthz
  ```
  Returns OK if the server is running and healthy.
- Metrics endpoint:
  Visit http://localhost:8080/metrics in a browser or with curl to see Prometheus metrics.
- Access web interface:
  Open http://localhost:8080/ in your browser to view the minimal front-end UI (served from the assets/ folder).
