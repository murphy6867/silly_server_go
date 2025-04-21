package chirp

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

	if len(body) < 0 {
		return nil, errors.New("chirp is too long")
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
