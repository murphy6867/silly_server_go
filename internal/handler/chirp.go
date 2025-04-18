package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/murphy6867/server/internal/model"
	"github.com/murphy6867/server/internal/validator"
	"github.com/murphy6867/server/utils"
)

func (vcq *ValidateChirps) ValidateChirpHandler(w http.ResponseWriter, r *http.Request) {
	var chirp model.ValidateChirp

	if err := json.NewDecoder(r.Body).Decode(&chirp); err != nil {
		log.Printf("Error encoding parameter: %s\n", err)
		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
		return
	}

	cleaned, err := validator.ValidateChirp(chirp.Body)
	if err != nil {
		log.Println("Validation error:", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"cleaned_body": cleaned})

}
