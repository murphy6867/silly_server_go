package chirp

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/murphy6867/server/internal/database"
	"github.com/murphy6867/server/utils"
)

type Chirp struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Body      string
	UserID    uuid.UUID
}

const MaxChirpLength = 140
const ReplaceString = "****"

var profaneWords = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

func NewChirp(userId string, body string) (*Chirp, error) {
	if len(body) > MaxChirpLength {
		return nil, errors.New("chirp is too long")
	}

	if len(body) == 0 {
		return nil, errors.New("chirp is too short")
	}

	cleaned := utils.FilterWord(profaneWords, body, ReplaceString)
	uId, err := uuid.Parse(userId)
	if err != nil {
		return nil, errors.New("user id not found")
	}

	return &Chirp{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Body:      cleaned,
		UserID:    uId,
	}, nil
}

func GetChirps(ctx context.Context, data []database.Chirp) (*ResponseChirpsDTO, error) {
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

	return &chirps, nil
}
