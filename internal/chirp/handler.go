package chirp

import (
	"encoding/json"
	"net/http"

	"github.com/murphy6867/silly_server_go/utils"
)

type ChirpHandler struct {
	svc *ChirpService
}

func NewChirpHandler(svc *ChirpService) *ChirpHandler {
	return &ChirpHandler{svc: svc}
}

func (c *ChirpHandler) CreateChirpHandler(w http.ResponseWriter, r *http.Request) {
	var dataIn CreateChirpDTO
	if err := json.NewDecoder(r.Body).Decode(&dataIn); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	chirp, err := c.svc.CreateChirpService(r.Context(), dataIn)
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

func (c *ChirpHandler) GetAllChirpsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := c.svc.GetAllChirpsService(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, out)
}

func (c *ChirpHandler) GetChirpsByIdHandler(w http.ResponseWriter, r *http.Request) {
	chirpId := r.PathValue("chirpID")

	out, err := c.svc.GetChirpsByIdService(r.Context(), chirpId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, http.StatusOK, out)

}
