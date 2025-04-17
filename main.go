package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/murphy6867/server/utils"
)

type clientConfig struct {
	fileserverHits atomic.Int32
}

type validateChirps struct {
	Body string `json:"body"`
}

func (cfg *clientConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *clientConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d\n", cfg.fileserverHits.Load())))
}

func (cfg *clientConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hit counter reset to 0\n"))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *clientConfig) metricsAdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", cfg.fileserverHits.Load())))
}

func (cfg *clientConfig) resetAdminHandler(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hit counter reset to 0\n"))
}

func (vcq *validateChirps) validateChirp(w http.ResponseWriter, r *http.Request) {

	const MaxChirpLength = 140

	if err := json.NewDecoder(r.Body).Decode(vcq); err != nil {
		log.Printf("Error encoding parameter: %s\n", err)
		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
		return
	}

	if countLength := len(vcq.Body); countLength > MaxChirpLength {
		log.Printf("Error Chirp is too long\n")
		utils.WriteJSON(w, 400, map[string]string{"error": "Chirp is too long"})
		return
	}

	filterProfane := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}

	response := utils.FilterWord(filterProfane, vcq.Body, "****")

	utils.WriteJSON(w, 200, map[string]string{"cleaned_body": response})
}

func main() {
	mux := http.NewServeMux()
	clientCfg := &clientConfig{}
	validateChirpReq := &validateChirps{}

	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/app/", clientCfg.middlewareMetricsInc(http.StripPrefix("/app", fileServer)))

	mux.HandleFunc("GET /api/healthz", healthCheck)
	mux.Handle("GET /api/assets/", http.StripPrefix("/api/assets/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("GET /api/metrics", clientCfg.metricsHandler)
	mux.HandleFunc("POST /api/reset", clientCfg.resetHandler)

	mux.HandleFunc("GET /admin/metrics", clientCfg.metricsAdminHandler)
	mux.HandleFunc("POST /admin/reset", clientCfg.resetAdminHandler)

	mux.HandleFunc("POST /api/validate_chirp", validateChirpReq.validateChirp)

	fmt.Println("Server start! Listening on :8080 🔥")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server error: ", err)
	}
}
