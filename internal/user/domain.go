package user

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
}

func NewUser(email string) (*User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("invalid email format")
	}

	return &User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
	}, nil
}
