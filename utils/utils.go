package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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

func FilterWord(profaneWords map[string]bool, sentence string, replaceString string) string {
	splitSentence := strings.Split(sentence, " ")

	for i, w := range splitSentence {
		if profaneWords[strings.ToLower(w)] {
			splitSentence[i] = replaceString
		}
	}
	return strings.Join(splitSentence, " ")
}
