package chirp

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/murphy6867/server/internal/database"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, c *Chirp) error
	GetAllChirps(ctx context.Context) ([]database.Chirp, error)
	GetChirpsById(ctx context.Context, userId uuid.UUID) (database.Chirp, error)
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

func (r *repository) GetAllChirps(ctx context.Context) ([]database.Chirp, error) {
	dbChirps, err := r.queries.GetAllChirp(ctx)
	if err != nil {
		return nil, err
	}
	return dbChirps, nil
}

func (r *repository) GetChirpsById(ctx context.Context, chirpId uuid.UUID) (database.Chirp, error) {
	dbChirps, err := r.queries.GetChirpById(ctx, chirpId)
	if err != nil {
		return dbChirps, err
	}
	return dbChirps, nil
}
