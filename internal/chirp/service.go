package chirp

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type ChirpService struct {
	repo ChirpRepository
}

func NewChirpService(c ChirpRepository) *ChirpService {
	return &ChirpService{repo: c}
}

func (svc *ChirpService) CreateChirpService(r *http.Request, data CreateChirpDTO) (*Chirp, error) {
	if data.Body == "" {
		return nil, errors.New("chirp is required")
	}

	chirp, err := NewChirp(r, data)
	if err != nil {
		return nil, err
	}

	chirp, err = svc.repo.CreateChirp(r.Context(), chirp)
	if err != nil {
		return nil, err
	}

	return chirp, nil
}

func (svc *ChirpService) GetAllChirpsService(ctx context.Context) (*ResponseChirpsDTO, error) {
	dbChirps, err := svc.repo.GetAllChirps(ctx)
	if err != nil {
		return nil, err
	}

	chirps := GetChirps(ctx, dbChirps)

	return chirps, nil
}

func (svc *ChirpService) GetChirpsByIdService(ctx context.Context, chirpId string) (*ResponseCreateChirpDTO, error) {
	if chirpId == "" {
		return nil, errors.New("parameter value is required")
	}
	parsedUserId, err := uuid.Parse(chirpId)
	if err != nil {
		return nil, err
	}

	dbChirps, err := svc.repo.GetChirpsById(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}

	chirps := GetChirpById(ctx, dbChirps)

	return chirps, nil

}
