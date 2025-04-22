package auth

import (
	"context"
	"errors"
)

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(r AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (u *AuthService) SignUpUserService(ctx context.Context, data SignUpUserDTO) (*User, error) {
	if data.Email == "" {
		return nil, errors.New("email is required")
	}
	if data.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := NewUser(data.Email, data.Password)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Register(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *AuthService) SignInService(ctx context.Context, data SignInDTO) (*User, error) {
	if data.Email == "" {
		return nil, errors.New("email is required")
	}
	if data.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := u.repo.SignIn(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	res, err := NewSignIn(user, data.Password)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return res, nil
}
