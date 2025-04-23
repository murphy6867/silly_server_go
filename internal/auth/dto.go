package auth

import (
	"time"

	"github.com/google/uuid"
)

type Option func(i SignInDTO) SignInDTO

type SignUpUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`

	ExpiresInSecond int32 `json:"expires_in_seconds"`
}

type ResponseUerDTO struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Email       string `json:"email"`
	AccessToken string `json:"token"`
}

type SignIn struct {
	Email           string
	Password        string
	ExpiresInSecond time.Duration
	AccessToken     string
}

type SignUpUserInfo struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
}

type User struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string
	AccessToken string
}
