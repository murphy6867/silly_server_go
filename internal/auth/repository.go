package auth

import (
	"context"
	"database/sql"

	"github.com/murphy6867/silly_server_go/internal/database"
)

type AuthRepository interface {
	Register(ctx context.Context, u *User) error
	SignIn(ctx context.Context, email string) (User, error)
}

type repository struct {
	queries *database.Queries
}

func NewRepository(db *sql.DB) AuthRepository {
	return &repository{
		queries: database.New(db),
	}
}

func (r *repository) Register(ctx context.Context, u *User) error {
	_, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		ID:             u.ID,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		Email:          u.Email,
		HashedPassword: u.Password,
	})
	return err
}

func (r *repository) SignIn(ctx context.Context, email string) (User, error) {
	u, err := r.queries.GetUserById(ctx, email)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
		Password:  u.HashedPassword,
	}, nil
}
