package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
}
type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseCreateUerDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

type ResponseUser struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
}
