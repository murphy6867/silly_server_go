package chirp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/database"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type ChirpService struct {
	repo ChirpRepository
}

func NewChirpService(c ChirpRepository) *ChirpService {
	return &ChirpService{repo: c}
}

func (svc *ChirpService) CreateChirpService(r *http.Request, data CreateChirpDTO) (*database.Chirp, error) {
	if data.Body == "" {
		return nil, utils.NewDomainError(400, "bad request")
	}

	chirp, err := NewChirp(r, data)
	if err != nil {
		return nil, err
	}

	return svc.repo.CreateChirp(r.Context(), chirp)
}

func (svc *ChirpService) GetAllChirpsService(ctx context.Context, sorted *string) (*[]database.Chirp, error) {
	if sorted == nil {
		defaultSort := "desc"
		sorted = &defaultSort
	}

	data, err := MappingChirpInfo("", false, sorted)
	if err != nil {
		return nil, err
	}

	dbChirps, err := svc.repo.GetAllChirps(ctx, data)
	if err != nil {
		return nil, err
	}

	return dbChirps, nil
}

func (svc *ChirpService) GetChirpsByUserIdService(ctx context.Context, userID, sorted *string) (*[]database.Chirp, error) {
	if userID == nil {
		return nil, errors.New("parameter value is required")
	}

	chirps, err := MappingChirpInfo(*userID, true, sorted)
	if err != nil {
		return nil, err
	}

	return svc.repo.GetChirpsByUserId(ctx, chirps)
}

func (svc *ChirpService) GetChirpsByIdService(ctx context.Context, chirpIdString *string) (*database.Chirp, error) {
	chirpID, err := uuid.Parse(*chirpIdString)
	if err != nil {
		return nil, utils.NewDomainError(400, fmt.Sprintf("bad request: %s", err))
	}

	return svc.repo.GetChirpById(ctx, chirpID)
}

func (svc *ChirpService) DeleteChirpByIdService(ctx context.Context, header http.Header, chirpId string) error {
	secTK := svc.repo.GetSecretKeyString()

	body, err := MappingChirpAndAuthorization(header, secTK, chirpId)
	if err != nil {
		return err
	}

	return svc.repo.DeleteChirpById(ctx, body)
}
