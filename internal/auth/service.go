package auth

import (
	"context"
	"errors"
	"net/http"
)

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(r AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) SignUpUserService(ctx context.Context, data SignUpUserDTO) (*SignUpUserInfo, error) {
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

	if err := s.repo.Register(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) SignInService(ctx context.Context, data SignInDTO) (*SignInResponse, error) {
	if data.Email == "" {
		return nil, errors.New("email is required")
	}

	if data.Password == "" {
		return nil, errors.New("password is required")
	}

	dataSignIn, err := NewSignIn(data)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	user, err := s.repo.SignIn(ctx, dataSignIn)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) RefreshTokenService(ctx context.Context, header http.Header) (*UserRefreshToken, error) {
	refreshToken, err := GetRefreshToken(header)
	if err != nil {
		return nil, err

	}
	return s.repo.RefreshTokenRepo(ctx, refreshToken)
}

func (s *AuthService) RevokeRefreshTokenService(ctx context.Context, header http.Header) error {
	refreshToken, err := GetRefreshToken(header)
	if err != nil {
		return err
	}
	return s.repo.RevokeRefreshTokenRepo(ctx, refreshToken)
}

func (s *AuthService) UpdateEmailAndPasswordService(ctx context.Context, header http.Header, body EditEmailAndPasswordDTO) (*SignInResponse, error) {
	mappingUser := EditEmailAndPassword(body)

	user, err := s.repo.UpdateEmailAndPasswordRepo(ctx, header, mappingUser)
	if err != nil {
		return nil, err
	}
	return user, err
}
