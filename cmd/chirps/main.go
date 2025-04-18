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
)

// type apiConfig struct {
// 	fileserverHits atomic.Int32
// 	DB             *database.Queries
// }

// type validateChirps struct {
// 	Body string `json:"body"`
// }

// func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Add(1)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Hits: %d\n", cfg.fileserverHits.Load())))
// }

// func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
// 	cfg.fileserverHits.Store(0)
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hit counter reset to 0\n"))
// }

// func healthCheck(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "Plain/text; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("OK"))
// }

// func (cfg *apiConfig) metricsAdminHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(http.StatusOK)

// 	w.Write([]byte(fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", cfg.fileserverHits.Load())))
// }

// func (cfg *apiConfig) resetAdminHandler(w http.ResponseWriter, r *http.Request) {
// 	cfg.fileserverHits.Store(0)
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hit counter reset to 0\n"))
// }

// func (vcq *validateChirps) validateChirp(w http.ResponseWriter, r *http.Request) {

// 	const MaxChirpLength = 140

// 	if err := json.NewDecoder(r.Body).Decode(vcq); err != nil {
// 		log.Printf("Error encoding parameter: %s\n", err)
// 		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
// 		return
// 	}

// 	if countLength := len(vcq.Body); countLength > MaxChirpLength {
// 		log.Printf("Error Chirp is too long\n")
// 		utils.WriteJSON(w, 400, map[string]string{"error": "Chirp is too long"})
// 		return
// 	}

// 	filterProfane := map[string]bool{
// 		"kerfuffle": true,
// 		"sharbert":  true,
// 		"fornax":    true,
// 	}

// 	response := utils.FilterWord(filterProfane, vcq.Body, "****")

// 	utils.WriteJSON(w, 200, map[string]string{"cleaned_body": response})
// }

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Database connection error: %v\n", err)
	}

	dbQueries := database.New(db)

	apiCfg := handler.APIConfig{
		DB: dbQueries,
	}
	validateChirpReq := handler.ValidateChirps{}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/app/", apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", fs)))
	mux.Handle("GET /api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir("../../web/static/images"))))

	// User
	mux.HandleFunc("GET /api/healthz", handler.HealthCheck)
	mux.HandleFunc("GET /api/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /api/reset", apiCfg.ResetHandler)

	// Admin
	mux.HandleFunc("GET /admin/metrics", apiCfg.MetricsHandler)
	mux.HandleFunc("POST /admin/reset", apiCfg.ResetHandler)

	mux.HandleFunc("POST /api/validate_chirp", validateChirpReq.ValidateChirpHandler)

	wrapped := apiCfg.MiddlewareMetricsInc(mux)

	fmt.Println("Server start! Listening on :8080")
	if err := http.ListenAndServe(":8080", wrapped); err != nil {
		fmt.Println("Server error: ", err)
	}
}
