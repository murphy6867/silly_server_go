package auth

import (
	"errors"
	"net/mail"
	"time"

	"github.com/google/uuid"
)

func NewUser(email string, password string) (*SignUpUserInfo, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("invalid email format")
	}

	hashedPas, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &SignUpUserInfo{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     email,
		Password:  hashedPas,
	}, nil
}

func NewSignIn(data SignInDTO) (*SignIn, error) {
	var expiresIn time.Duration
	if data.ExpiresInSecond == 0 {
		expiresIn = time.Hour
	} else if data.ExpiresInSecond > 3600 {
		expiresIn = time.Hour
	} else {
		expiresIn = time.Duration(data.ExpiresInSecond) * time.Second
	}

	return &SignIn{
		Email:           data.Email,
		Password:        data.Password,
		ExpiresInSecond: expiresIn,
	}, nil
}
