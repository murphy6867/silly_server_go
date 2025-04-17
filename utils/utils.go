package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data, err := json.Marshal(payload); err == nil {
		w.Write(data)
	} else {
		log.Printf("Error marshalling response: %v", err)
	}
}
