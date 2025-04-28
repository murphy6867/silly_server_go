package webhook

import (
	"context"
	"database/sql"

	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type WebHookRepository interface {
	UpdateSubscriptionChirpyRed(ctx context.Context, body *User) error
	GetAPIKeyString() *string
}

type repository struct {
	queries   *database.Queries
	secretKey string
	PolkaKey  string
}

func NewRepository(queries *database.Queries, secretKey, PolkaKey string) WebHookRepository {
	return &repository{
		queries:   queries,
		secretKey: secretKey,
		PolkaKey:  PolkaKey,
	}
}

func (r *repository) UpdateSubscriptionChirpyRed(ctx context.Context, body *User) error {
	err := r.queries.UpdateIsChirpsRedStatus(ctx, database.UpdateIsChirpsRedStatusParams{
		ID:          body.ID,
		IsChirpyRed: sql.NullBool{Bool: body.IsChirpyRed, Valid: true},
	})
	if err != nil {
		return utils.NewDomainError(404, "user not found")
	}

	return nil
}

func (r *repository) GetAPIKeyString() *string {
	return &r.PolkaKey
}
