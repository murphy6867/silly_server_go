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
		utils.HandleError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, mapDatabaseChirp(*chirp))
}

func (c *ChirpHandler) GetChirpsHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("author_id")
	chirpID := q.Get("id")
	sort := q.Get("sort")
	sortPtr := &sort

	switch {
	case chirpID != "" && userID == "":
		chirp, err := c.svc.GetChirpsByIdService(r.Context(), &chirpID)
		if err != nil {
			utils.HandleError(w, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, mapDatabaseChirp(*chirp))

	case userID != "" && chirpID == "":
		userDto, err := c.svc.GetChirpsByUserIdService(r.Context(), &userID, sortPtr)
		if err != nil {
			utils.HandleError(w, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, mapDatabaseChirps(*userDto))

	case userID == "" && chirpID == "":
		domainChirps, err := c.svc.GetAllChirpsService(r.Context(), sortPtr)
		if err != nil {
			utils.HandleError(w, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, mapDatabaseChirps(*domainChirps))

	default:
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
	}
}

func (c *ChirpHandler) DeleteChirpByIdHandler(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")

	if err := c.svc.DeleteChirpByIdService(r.Context(), r.Header, chirpIDString); err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
