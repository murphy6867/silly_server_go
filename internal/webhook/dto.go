package webhook

import (
	"time"

	"github.com/google/uuid"
)

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
