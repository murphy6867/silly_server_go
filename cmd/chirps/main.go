package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/murphy6867/server/internal/database"
	"github.com/murphy6867/server/internal/handler"
	"github.com/murphy6867/server/internal/user"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	apiCfg := handler.APIConfig{DB: dbQueries}
	validateChirpReq := handler.ValidateChirps{}

	// Compose User module
	userRepo := user.NewRepository(db)
	userSvc := user.NewUserService(userRepo)
	useHdl := user.NewUserHandler(userSvc)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", fs)))
	mux.Handle("GET /api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir("../../web/static/images"))))

	// Public API routes
	mux.HandleFunc("GET /api/healthz", handler.HealthCheck)
	mux.HandleFunc("GET /api/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /api/reset", apiCfg.ResetHandler)
	mux.HandleFunc("POST /api/validate_chirp", validateChirpReq.ValidateChirpHandler)
	mux.HandleFunc("POST /api/users", useHdl.CreateUserHandler)
	// mux.HandleFunc("POST /api/chirps", useHdl.CreateUserHandler)

	// Admin API routes
	mux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)

	wrappedMux := apiCfg.MiddlewareMetricsInc(mux)

	addr := ":8080"
	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
