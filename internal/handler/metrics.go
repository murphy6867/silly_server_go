package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *APIConfig) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d\n", cfg.FileServerHits.Load())))
}

func (cfg *APIConfig) ResetHandler(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	platForm := os.Getenv("PLATFORM")
	if platForm != "dev" {
		utils.WriteJSON(w, http.StatusForbidden, map[string]string{"error": "Access denied"})
	}
	cfg.FileServerHits.Store(0)

	if err := cfg.DB.ResetUserTable(r.Context()); err != nil {
		log.Printf("Error reset fail: %s\n", err)
		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
		return
	}

	w.WriteHeader(http.StatusOK)
}
