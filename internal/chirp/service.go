package chirp

import (
	"context"
	"errors"
)

type ChirpService struct {
	repo ChirpRepository
}

func NewChirpService(c ChirpRepository) *ChirpService {
	return &ChirpService{repo: c}
}

func (c *ChirpService) CreateChirpService(ctx context.Context, data CreateChirpDTO) (*Chirp, error) {
	if data.Body == "" {
		return nil, errors.New("chirp is required")
	}

	if data.UserID == "" {
		return nil, errors.New("user id is required")
	}

	chirp, err := NewChirp(data.UserID, data.Body)
	if err != nil {
		return nil, err
	}

	if err := c.repo.CreateChirp(ctx, chirp); err != nil {
		return nil, err
	}

	return chirp, nil
}

func (c *ChirpService) GetAllChirpsService(ctx context.Context) (*ResponseChirpsDTO, error) {
	dbChirps, err := c.repo.GetAllChirps(ctx)
	if err != nil {
		return nil, err
	}

	chirps, err := GetChirps(ctx, dbChirps)

	return chirps, nil
}
