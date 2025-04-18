package handler

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
