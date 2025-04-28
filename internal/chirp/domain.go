package chirp

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/murphy6867/silly_server_go/internal/auth"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

const MaxChirpLength = 140
const ReplaceString = "****"

var profaneWords = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

func NewChirp(r *http.Request, data CreateChirpDTO) (*Chirp, error) {
	token, err := utils.GetBearerToken(r.Header)
	if err != nil {
		return nil, utils.NewDomainError(401, fmt.Sprintf("unauthorized: %s", err))

	}

	if len(data.Body) > MaxChirpLength {
		return nil, utils.NewDomainError(400, "chirp is too long")

	}

	if len(data.Body) == 0 {
		return nil, utils.NewDomainError(400, "chirp is too short")

	}

	cleaned := utils.FilterWord(profaneWords, data.Body, ReplaceString)

	return &Chirp{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Body:        cleaned,
		AccessToken: token,
	}, nil
}

func MappingChirpInfo(stringID string, isUserID bool, sortedString *string) (*ManageChirpInfo, error) {
	info := &ManageChirpInfo{}

	info.SortString = *sortedString
	if stringID == "" {
		return info, nil
	}

	parsedID, err := uuid.Parse(stringID)
	if err != nil {
		return nil, utils.NewDomainError(400, fmt.Sprintf("bad request: %s", err))
	}

	if isUserID {
		info.UserID = parsedID
	} else {
		info.ChirpID = parsedID
	}

	return info, nil
}

func MappingChirpAndAuthorization(header http.Header, secretKey string, chirpId string) (*ManageChirpInfo, error) {
	token, err := utils.GetBearerToken(header)
	if err != nil {
		return nil, utils.NewDomainError(401, "token not found")
	}

	userID, err := auth.ValidateJWT(token, secretKey)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	chirpID, err := uuid.Parse(chirpId)
	if err != nil {
		return nil, errors.New("invalid chirp id")
	}

	return &ManageChirpInfo{
		UserID:      userID,
		AccessToken: token,
		ChirpID:     chirpID,
	}, nil
}
