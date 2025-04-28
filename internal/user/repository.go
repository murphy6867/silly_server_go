package user

import (
	"github.com/murphy6867/silly_server_go/internal/database"
)

type UserRepository interface {
	// UpdateSubscriptionChirpyRed(ctx context.Context, body *User) error
}

type repository struct {
	queries   *database.Queries
	secretKey string
}

func NewRepository(queries *database.Queries, secretKey string) UserRepository {
	return &repository{
		queries:   queries,
		secretKey: secretKey,
	}
}

// func (r *repository) UpdateSubscriptionChirpyRed2(ctx context.Context, body *User) error {
// 	err := r.queries.UpdateIsChirpsRedStatus(ctx, database.UpdateIsChirpsRedStatusParams{
// 		ID:          body.ID,
// 		IsChirpyRed: sql.NullBool{Bool: body.IsChirpyRed, Valid: true},
// 	})
// 	if err != nil {
// 		return utils.NewDomainError(404, "user not found")
// 	}

// 	return nil
// }
