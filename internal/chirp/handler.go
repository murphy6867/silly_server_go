package chirp

import (
	"encoding/json"
	"net/http"

	"github.com/murphy6867/server/utils"
)

type ChirpHandler struct {
	svc *ChirpService
}

func NewChirpHandler(svc *ChirpService) *ChirpHandler {
	return &ChirpHandler{svc: svc}
}

func (h *ChirpHandler) CreateChirpHandler(w http.ResponseWriter, r *http.Request) {
	var dataIn CreateChirpDTO
	if err := json.NewDecoder(r.Body).Decode(&dataIn); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	chirp, err := h.svc.CreateChirpService(r.Context(), dataIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := ResponseCreateChirpDTO{
		ID:        chirp.ID.String(),
		CreatedAt: chirp.CreatedAt.String(),
		UpdatedAt: chirp.UpdatedAt.String(),
		Body:      chirp.Body,
		UserID:    chirp.UserID.String(),
	}

	utils.WriteJSON(w, http.StatusCreated, out)
}
