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

	user, err := NewUser(data.Email)
	if err != nil {
		return nil, err
	}

	if err := u.repo.Register(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
