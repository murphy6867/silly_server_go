package user

import (
	"context"

	"github.com/murphy6867/silly_server_go/internal/database"
)

type UserRepository interface {
	Register(ctx context.Context, u *User) error
	SignIn(ctx context.Context, email string) (User, error)
}

type repository struct {
	queries *database.Queries
}
