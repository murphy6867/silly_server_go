package chirp

import (
	"encoding/json"
	"net/http"

	utils "github.com/murphy6867/silly_server_go/internal/shared"
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

	chirp, err := c.svc.CreateChirpService(r, dataIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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

func (c *ChirpHandler) GetChirpByIdHandler(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")

	out, err := c.svc.GetChirpsByIdService(r.Context(), chirpIDString)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &ResponseCreateChirpDTO{
		ID:        out.ID.String(),
		UserID:    out.UserID.String(),
		CreatedAt: out.CreatedAt.String(),
		UpdatedAt: out.UpdatedAt.String(),
		Body:      out.Body,
	})
}

func (c *ChirpHandler) DeleteChirpByIdHandler(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")

	if err := c.svc.DeleteChirpByIdService(r.Context(), r.Header, chirpIDString); err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
