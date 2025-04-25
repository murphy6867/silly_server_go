package auth

import (
	"errors"
	"net/http"
	"net/mail"
	"time"

	"github.com/google/uuid"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
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

func NewSignIn(data SignInDTO) (*SignInUserInfo, error) {
	refToken, err := MakeRefreshToken()
	if err != nil {
		return nil, err
	}

	return &SignInUserInfo{
		Email:             data.Email,
		Password:          data.Password,
		AccessTokenExpAt:  time.Hour,
		RefreshToken:      refToken,
		RefreshTokenExpAt: time.Hour * 24 * 60,
	}, nil
}

func GetRefreshToken(header http.Header) (*UserRefreshToken, error) {
	refreshTK, err := utils.GetBearerToken(header)
	if err != nil {
		return nil, err
	}

	return &UserRefreshToken{
		RefreshToken: refreshTK,
	}, nil
}
