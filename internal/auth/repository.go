package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type AuthRepository interface {
	Register(ctx context.Context, u *SignUpUserInfo) error
	SignIn(ctx context.Context, data *SignInUserInfo) (*SignInResponse, error)
	RefreshTokenRepo(ctx context.Context, refTK *UserRefreshToken) (*UserRefreshToken, error)
	RevokeRefreshTokenRepo(ctx context.Context, refTK *UserRefreshToken) error
	UpdateEmailAndPasswordRepo(ctx context.Context, header http.Header, body EditEmailAndPassword) (*SignInResponse, error)
	GetSecretKeyString() string
}

type repository struct {
	queries   *database.Queries
	secretKey string
}

func NewRepository(queries *database.Queries, secretKey string) AuthRepository {
	return &repository{
		queries:   queries,
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

func (r *repository) SignIn(ctx context.Context, data *SignInUserInfo) (*SignInResponse, error) {
	user, err := r.queries.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if err := CheckPasswordHash(data.Password, user.HashedPassword); err != nil {
		fmt.Println(err)
		return nil, errors.New("password incorrect")
	}

	accessToken, err := MakeJWT(user.ID, r.secretKey, data.AccessTokenExpAt)
	if err != nil {
		return nil, err
	}

	_, err = r.queries.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     data.RefreshToken,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		User: User{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token:        accessToken,
		RefreshToken: data.RefreshToken,
	}, nil
}

func (r *repository) RefreshTokenRepo(ctx context.Context, refTK *UserRefreshToken) (*UserRefreshToken, error) {
	user, err := r.queries.GetUserFromRefreshToken(ctx, refTK.RefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, err := MakeJWT(
		user.ID,
		r.secretKey,
		time.Hour,
	)
	if err != nil {
		return nil, err
	}

	return &UserRefreshToken{
		RefreshToken: accessToken,
	}, nil
}

func (r *repository) RevokeRefreshTokenRepo(ctx context.Context, refTK *UserRefreshToken) error {
	_, err := r.queries.RevokeRefreshToken(ctx, refTK.RefreshToken)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateEmailAndPasswordRepo(ctx context.Context, header http.Header, body EditEmailAndPassword) (*SignInResponse, error) {
	accessToken, err := utils.GetBearerToken(header)
	if err != nil {
		return nil, err
	}

	userId, err := ValidateJWT(accessToken, r.secretKey)
	if err != nil {
		return nil, err
	}

	newPass, err := HashPassword(body.NewPassword)
	if err != nil {
		return nil, err
	}

	newTime := time.Now().UTC()

	if err := r.queries.UpdateEmailAndPassword(ctx, database.UpdateEmailAndPasswordParams{
		ID:             userId,
		Email:          body.NewEmail,
		HashedPassword: newPass,
		UpdatedAt:      newTime,
	}); err != nil {
		return nil, err
	}

	return &SignInResponse{
		User: User{
			ID:        userId,
			Email:     body.NewEmail,
			CreatedAt: newTime,
		},
		RefreshToken: accessToken,
	}, nil
}

func (r *repository) GetSecretKeyString() string {
	return r.secretKey
}
