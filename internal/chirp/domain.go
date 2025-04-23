package chirp

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func GetChirpById(ctx context.Context, data database.Chirp) *ResponseCreateChirpDTO {
	return &ResponseCreateChirpDTO{
		ID:        data.ID.String(),
		CreatedAt: data.CreatedAt.Format(time.RFC3339),
		UpdatedAt: data.UpdatedAt.Format(time.RFC3339),
		UserID:    data.UserID.String(),
		Body:      data.Body,
	}
}
