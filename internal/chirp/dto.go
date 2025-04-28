package chirp

import (
	"time"

	"github.com/google/uuid"
)

type CreateChirpDTO struct {
	UserID      string `json:"user_id"`
	Body        string `json:"body"`
	AccessToken string `json:"token"`
}

type ResponseCreateChirpDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
}

type ChirpDTO struct {
	ChirpId string `json:"chirp_id"`
}

type Chirp struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Body        string
	UserID      uuid.UUID
	AccessToken string
}

type ManageChirpInfo struct {
	UserID      uuid.UUID
	AccessToken string
	ChirpID     uuid.UUID
	SortString  string
}

type ResponseChirpsDTO []ResponseCreateChirpDTO
