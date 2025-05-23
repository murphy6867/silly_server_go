package auth

import (
	"time"

	"github.com/google/uuid"
)

type SignUpUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditEmailAndPasswordDTO struct {
	NewEmail    string `json:"email"`
	NewPassword string `json:"password"`
}

type User struct {
	ID             uuid.UUID `json:"id,omitempty"`
	Email          string    `json:"email,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
	IsChirpyRed    bool      `json:"is_chirpy_red"`
}
type SignInResponse struct {
	User         User   `json:"user,omitempty"`
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	IsChirpyRed  bool   `json:"is_chirpy_red"`
	Email        string `json:"email,omitempty"`
}

type ResponseUerDTO struct {
	ID           string `json:"id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Email        string `json:"email"`
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	IsChirpyRed  bool   `json:"is_chirpy_red"`
}

type RefreshResponse struct {
	Token string `json:"token"`
}

type SignInUserInfo struct {
	Email             string
	Password          string
	AccessToken       string
	RefreshToken      string
	AccessTokenExpAt  time.Duration
	RefreshTokenExpAt time.Duration
}

type ResponseUserSignIn struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
}

type SignUpUserInfo struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string
	Password    string
	IsChirpyRed bool
}

type UserRefreshToken struct {
	RefreshToken string
}

type UserAccessToken struct {
	AccessToken string
}

type EditEmailAndPassword struct {
	NewEmail    string
	NewPassword string
}
