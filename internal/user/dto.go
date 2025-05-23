package user

import (
	"time"

	"github.com/google/uuid"
)

type RequestUerDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

type Data struct {
	UserID string `json:"user_id"`
}

type UpdateSubscriptionDTO struct {
	Data  Data   `json:"data"`
	Event string `json:"event"`
}

type User struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string
	IsChirpyRed bool
}
