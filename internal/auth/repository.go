package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/murphy6867/silly_server_go/internal/database"
)

type AuthRepository interface {
	Register(ctx context.Context, u *SignUpUserInfo) error
	SignIn(ctx context.Context, data *SignIn) (User, error)
}

type repository struct {
	queries   *database.Queries
	secretKey string
}

func NewRepository(db *sql.DB, secretKey string) AuthRepository {
	return &repository{
		queries:   database.New(db),
		secretKey: secretKey,
	}
}

func (r *repository) Register(ctx context.Context, u *SignUpUserInfo) error {
	_, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		ID:             u.ID,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
		Email:          u.Email,
		HashedPassword: u.Password,
	})
	return err
}

func (r *repository) SignIn(ctx context.Context, data *SignIn) (User, error) {
	u, err := r.queries.GetUserById(ctx, data.Email)
	if err != nil {
		return User{}, err
	}

	if err := CheckPasswordHash(data.Password, u.HashedPassword); err != nil {
		fmt.Println(err)
		return User{}, errors.New("password incorrect")
	}
	genToken, err := MakeJWT(u.ID, r.secretKey, data.ExpiresInSecond)
	if err != nil {
		return User{}, err
	}

	userToken, err := r.queries.SetUserToken(ctx, database.SetUserTokenParams{
		AccessToken: genToken,
		Email:       data.Email,
	})
	if err != nil {
		return User{}, err
	}

	return User{
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		Email:       u.Email,
		AccessToken: userToken.AccessToken,
	}, nil
}
