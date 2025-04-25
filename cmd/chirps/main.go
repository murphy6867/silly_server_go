package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/murphy6867/silly_server_go/internal"
	"github.com/murphy6867/silly_server_go/internal/auth"
	"github.com/murphy6867/silly_server_go/internal/chirp"
	"github.com/murphy6867/silly_server_go/internal/infra"
	"github.com/murphy6867/silly_server_go/internal/infra/config"
	"github.com/murphy6867/silly_server_go/internal/middleware"
)

func main() {
	apiCfg := config.Load()
	defer apiCfg.CloseDB.Close()

	// Compose Auth module
	authRepo := auth.NewRepository(apiCfg.DB, apiCfg.JWTSecret)
	authSvc := auth.NewAuthService(authRepo)
	authHld := auth.NewAuthHandler(authSvc)

	// Compose Chirp module
	chirpRepo := chirp.NewRepository(apiCfg.DB, apiCfg.JWTSecret)
	chirpSvc := chirp.NewChirpService(chirpRepo)
	chirpHdl := chirp.NewChirpHandler(chirpSvc)

	// Compose User module

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/app/", middleware.MiddlewareMetricsInc(http.StripPrefix("/app", fs), &apiCfg.FileServerHits))
	mux.Handle("GET /api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir("../../web/static/images"))))

	// Public API routes
	mux.HandleFunc("GET /api/healthz", internal.HealthCheck)
	// Auth
	mux.HandleFunc("POST /api/signin", authHld.SignInHandler)
	mux.HandleFunc("POST /api/login", authHld.SignInHandler)
	mux.HandleFunc("POST /api/signup", authHld.SignUpHandler)
	mux.HandleFunc("POST /api/users", authHld.SignUpHandler)
	mux.HandleFunc("POST /api/refresh", authHld.RefreshTokenHandler)
	mux.HandleFunc("POST /api/revoke", authHld.RevokeRefreshToken)
	// Chirp
	mux.HandleFunc("GET /api/chirps", chirpHdl.GetAllChirpsHandler)
	mux.HandleFunc("POST /api/chirps", chirpHdl.CreateChirpHandler)
	mux.HandleFunc("/api/chirps/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			chirpHdl.GetChirpsByIdHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// Admin API routes
	// TODO: Create internal/admin
	mux.HandleFunc("GET /admin/metrics", func(w http.ResponseWriter, r *http.Request) {
		infra.MetricsHandler(w, r, &apiCfg.FileServerHits)
	})
	mux.HandleFunc("POST /admin/reset", func(w http.ResponseWriter, r *http.Request) {
		infra.ResetHandler(w, r, &apiCfg.FileServerHits, apiCfg.DB)
	})

	wrappedMux := middleware.MiddlewareMetricsInc(mux, &apiCfg.FileServerHits)

	addr := ":8080"
	fmt.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}
