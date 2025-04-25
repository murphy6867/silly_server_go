package infra

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request, fsh *atomic.Int32) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d\n", fsh.Load())))
}

func ResetHandler(w http.ResponseWriter, r *http.Request, fsh *atomic.Int32, DB *database.Queries) {
	godotenv.Load()
	platForm := os.Getenv("PLATFORM")
	if platForm != "dev" {
		utils.WriteJSON(w, http.StatusForbidden, map[string]string{"error": "Access denied"})
	}
	fsh.Store(0)

	if err := DB.ResetUserTable(r.Context()); err != nil {
		log.Printf("Error reset fail: %s\n", err)
		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
		return
	}

	w.WriteHeader(http.StatusOK)
}
