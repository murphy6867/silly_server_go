package user

import (
	"context"
	"database/sql"

	"github.com/murphy6867/silly_server_go/internal/database"
)

type UserRepository interface {
	Register(ctx context.Context, u *User) error
}

type repository struct {
	queries *database.Queries
}

func NewRepository(db *sql.DB) UserRepository {
	return &repository{
		queries: database.New(db),
	}
}

func (r *repository) Register(ctx context.Context, u *User) error {
	_, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
	})
	return err
}
