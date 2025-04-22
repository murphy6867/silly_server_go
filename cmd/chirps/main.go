package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/murphy6867/server/internal/chirp"
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

	// Compose User module
	userRepo := user.NewRepository(db)
	userSvc := user.NewUserService(userRepo)
	useHdl := user.NewUserHandler(userSvc)

	// Compise Chirp module
	chirpRepo := chirp.NewRepository(db)
	chirpSvc := chirp.NewChirpService(chirpRepo)
	chirpHdl := chirp.NewChirpHandler(chirpSvc)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", fs)))
	mux.Handle("GET /api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir("../../web/static/images"))))

	// Public API routes
	mux.HandleFunc("GET /api/healthz", handler.HealthCheck)
	mux.HandleFunc("GET /api/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /api/reset", apiCfg.ResetHandler)
	// User
	mux.HandleFunc("POST /api/users", useHdl.CreateUserHandler)
	// Chirp
	mux.HandleFunc("GET /api/chirps/{chirpID}", chirpHdl.GetChirpsByIdHandler)
	mux.HandleFunc("GET /api/chirps", chirpHdl.GetAllChirpsHandler)
	mux.HandleFunc("POST /api/chirps", chirpHdl.CreateChirpHandler)

	// Admin API routes
	// TODO: Create internal/admin
	mux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)

	wrappedMux := apiCfg.MiddlewareMetricsInc(mux)

	addr := ":8080"
	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
