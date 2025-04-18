package handler

import (
	"log"
	"net/http"
)

func (cfg *APIConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requested path:", r.URL.Path)
		cfg.FileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
