package chirp

import (
	"context"
	"errors"
	"net/http"
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

func (svc *ChirpService) GetChirpsByIdService(ctx context.Context, chirpID string) (*Chirp, error) {
	if chirpID == "" {
		return nil, errors.New("parameter value is required")
	}

	chirps, err := MappingChirp(chirpID)
	if err != nil {
		return nil, err
	}

	dbChirps, err := svc.repo.GetChirpById(ctx, chirps.ChirpID)
	if err != nil {
		return nil, err
	}

	return dbChirps, nil
}

func (svc *ChirpService) DeleteChirpByIdService(ctx context.Context, header http.Header, chirpId string) error {
	secTK := svc.repo.GetSecretKeyString()

	body, err := MappingChirpAndAuthorization(header, secTK, chirpId)
	if err != nil {
		return err
	}

	return svc.repo.DeleteChirpById(ctx, body)
}
