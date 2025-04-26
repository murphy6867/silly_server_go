package chirp

import (
	"context"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/auth"
	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, c *Chirp) (*Chirp, error)
	GetAllChirps(ctx context.Context) ([]database.Chirp, error)
	GetChirpById(ctx context.Context, userId uuid.UUID) (*Chirp, error)
	DeleteChirpById(ctx context.Context, data *ManageChirpInfo) error
	GetSecretKeyString() string
}

type repository struct {
	queries   *database.Queries
	secretKey string
}

func NewRepository(queries *database.Queries, secretKey string) ChirpRepository {
	return &repository{
		queries:   queries,
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

func (r *repository) GetChirpById(ctx context.Context, chirpId uuid.UUID) (*Chirp, error) {
	dbChirps, err := r.queries.GetChirpById(ctx, chirpId)
	if err != nil {
		return nil, utils.NewDomainError(404, "not found")
	}
	return &Chirp{
		ID:        dbChirps.ID,
		UserID:    dbChirps.UserID,
		Body:      dbChirps.Body,
		CreatedAt: dbChirps.CreatedAt,
		UpdatedAt: dbChirps.UpdatedAt,
	}, nil
}

func (r *repository) DeleteChirpById(ctx context.Context, data *ManageChirpInfo) error {
	dbChirp, err := r.queries.GetChirpById(ctx, data.ChirpID)
	if err != nil {
		return utils.NewDomainError(404, "chirp not found")
	}

	if dbChirp.UserID != data.UserId {
		return utils.NewDomainError(403, "forbidden permission error")
	}

	err = r.queries.DeleteChirpById(ctx, dbChirp.ID)
	if err != nil {
		return utils.NewDomainError(500, "internal Server error")
	}
	return nil
}

func (r *repository) GetSecretKeyString() string {
	return r.secretKey
}
