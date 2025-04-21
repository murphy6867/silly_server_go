package chirp

import (
	"context"
	"database/sql"

	"github.com/murphy6867/server/internal/database"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, c *Chirp) error
}

type repository struct {
	queries *database.Queries
}

func NewRepository(db *sql.DB) ChirpRepository {
	return &repository{
		queries: database.New(db),
	}
}

func (r *repository) CreateChirp(ctx context.Context, c *Chirp) error {
	_, err := r.queries.CreateChirp(ctx, database.CreateChirpParams{
		ID:        c.ID,
		UserID:    c.UserID,
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	return err
}
