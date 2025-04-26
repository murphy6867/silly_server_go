package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

type DomainError struct {
	Code    int
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

func NewDomainError(code int, message string) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

func HandleError(w http.ResponseWriter, err error) {
	if domainErr, ok := err.(*DomainError); ok {
		WriteJSON(w, domainErr.Code, map[string]string{
			"error": domainErr.Message,
		})
		return
	}

	WriteJSON(w, http.StatusInternalServerError, map[string]string{
		"error": "Internal Server Error",
	})
}

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
	words := strings.Fields(sentence)

	for i, w := range words {
		if profaneWords[strings.ToLower(w)] {
			words[i] = replaceString
		}
	}
	return strings.Join(words, " ")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Plain/text; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get("Authorization")
	if token == "" {
		return "", errors.New("invalid authorization")
	}
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token, nil
}
