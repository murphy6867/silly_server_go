package user

import (
	"context"
	"errors"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

func (u *UserService) CreateUserService(ctx context.Context, data CreateUserDTO) (*User, error) {
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

func (u *UserService) SignInService(ctx context.Context, data SignInUserDTO) (*ResponseUser, error) {
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
