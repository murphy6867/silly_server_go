package chirp

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/auth"
	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

const MaxChirpLength = 140
const ReplaceString = "****"

var profaneWords = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

func NewChirp(r *http.Request, data CreateChirpDTO) (*Chirp, error) {
	token, err := utils.GetBearerToken(r.Header)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if len(data.Body) > MaxChirpLength {
		return nil, errors.New("chirp is too long")
	}

	if len(data.Body) == 0 {
		return nil, errors.New("chirp is too short")
	}

	cleaned := utils.FilterWord(profaneWords, data.Body, ReplaceString)

	return &Chirp{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Body:        cleaned,
		AccessToken: token,
	}, nil
}

func GetChirps(ctx context.Context, data []database.Chirp) *ResponseChirpsDTO {
	chirps := make(ResponseChirpsDTO, len(data))
	for i, ch := range data {
		chirps[i] = ResponseCreateChirpDTO{
			ID:        ch.ID.String(),
			UserID:    ch.UserID.String(),
			Body:      ch.Body,
			CreatedAt: ch.CreatedAt.Format(time.RFC3339),
			UpdatedAt: ch.UpdatedAt.Format(time.RFC3339),
		}
	}

	return &chirps
}

func MappingChirp(chirpID string) (*ManageChirpInfo, error) {
	parsedChirpID, err := uuid.Parse(chirpID)
	if err != nil {
		return nil, utils.NewDomainError(400, "chirp id notfound")
	}
	return &ManageChirpInfo{
		ChirpID: parsedChirpID,
	}, nil
}

func MappingChirpAndAuthorization(header http.Header, secretKey string, chirpId string) (*ManageChirpInfo, error) {
	token, err := utils.GetBearerToken(header)
	if err != nil {
		return nil, utils.NewDomainError(401, "token not found")
	}

	userID, err := auth.ValidateJWT(token, secretKey)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	chirpID, err := uuid.Parse(chirpId)
	if err != nil {
		return nil, errors.New("invalid chirp id")
	}

	return &ManageChirpInfo{
		UserId:      userID,
		AccessToken: token,
		ChirpID:     chirpID,
	}, nil
}
