package auth

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

func NewUser(email string, password string) (*User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("invalid email format")
	}

	hashedPas, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     email,
		Password:  hashedPas,
	}, nil
}

func NewSignIn(user User, password string) (*User, error) {
	if err := checkPasswordHash(user.Password, password); err != nil {
		return nil, errors.New("password incorrect")
	}

	return &User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
