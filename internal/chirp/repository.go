package chirp

import (
	"context"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/auth"
	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type ChirpRepository interface {
	CreateChirp(ctx context.Context, c *Chirp) (*database.Chirp, error)
	GetChirpById(ctx context.Context, chirpId uuid.UUID) (*database.Chirp, error)
	GetChirpsByUserId(ctx context.Context, data *ManageChirpInfo) (*[]database.Chirp, error)
	GetAllChirps(ctx context.Context, data *ManageChirpInfo) (*[]database.Chirp, error)
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

func (r *repository) CreateChirp(ctx context.Context, c *Chirp) (*database.Chirp, error) {
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

	return &data, err
}

func (r *repository) GetChirpById(ctx context.Context, chirpId uuid.UUID) (*database.Chirp, error) {
	dbChirps, err := r.queries.GetChirpById(ctx, chirpId)
	if err != nil {
		return nil, utils.NewDomainError(404, "not found")
	}
	return &dbChirps, nil
}

func (r *repository) GetAllChirps(
	ctx context.Context,
	info *ManageChirpInfo,
) (*[]database.Chirp, error) {
	order := strings.ToLower(info.SortString)

	var (
		chirps []database.Chirp
		err    error
	)
	if order == "desc" {
		chirps, err = r.queries.GetAllChirpsDesc(ctx)
		log.Println("========= DESC =====> ", chirps)
	} else {
		chirps, err = r.queries.GetAllChirpsAsc(ctx)
		log.Println("========= ASC =====> ", chirps)

	}
	if err != nil {
		return nil, err
	}

	return &chirps, nil
}

func (r *repository) GetChirpsByUserId(
	ctx context.Context,
	info *ManageChirpInfo,
) (*[]database.Chirp, error) {
	order := strings.ToLower(info.SortString)

	var (
		chirps []database.Chirp
		err    error
	)
	if order == "desc" {
		chirps, err = r.queries.GetChirpsByUserIdDesc(ctx, info.UserID)
	} else {
		chirps, err = r.queries.GetChirpsByUserIdAsc(ctx, info.UserID)
	}
	if err != nil {
		return nil, err
	}

	return &chirps, nil
}

func (r *repository) DeleteChirpById(ctx context.Context, data *ManageChirpInfo) error {
	dbChirp, err := r.queries.GetChirpById(ctx, data.ChirpID)
	if err != nil {
		return utils.NewDomainError(404, "chirp not found")
	}

	if dbChirp.UserID != data.UserID {
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
