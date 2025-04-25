package middleware

import (
	"log"
	"net/http"
	"sync/atomic"
)

func MiddlewareMetricsInc(next http.Handler, fsh *atomic.Int32) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requested path:", r.URL.Path)
		fsh.Add(1)
		next.ServeHTTP(w, r)
	})
}
