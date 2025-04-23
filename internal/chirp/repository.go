package chirp

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/auth"
	"github.com/murphy6867/silly_server_go/internal/database"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, c *Chirp) (*Chirp, error)
	GetAllChirps(ctx context.Context) ([]database.Chirp, error)
	GetChirpsById(ctx context.Context, userId uuid.UUID) (database.Chirp, error)
}

type repository struct {
	queries   *database.Queries
	secretKey string
}

func NewRepository(db *sql.DB, secretKey string) ChirpRepository {
	return &repository{
		queries:   database.New(db),
		secretKey: secretKey,
	}
}

func (r *repository) CreateChirp(ctx context.Context, c *Chirp) (*Chirp, error) {
	userID, err := auth.ValidateJWT(c.AccessToken, r.secretKey)
	if err != nil {
		return nil, err
	}

	data, err := r.queries.CreateChirp(ctx, database.CreateChirpParams{
		ID:        c.ID,
		UserID:    userID,
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})

	return &Chirp{
		ID:        data.ID,
		UserID:    userID,
		Body:      data.Body,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, err
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
