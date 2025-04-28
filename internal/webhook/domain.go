package webhook

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/murphy6867/silly_server_go/internal/auth"
	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

func MappingUpdateSubscription(header http.Header, body *UpdateSubscriptionDTO, apiKeyString *string) (*User, error) {
	validApiKey, err := auth.GetAPIKey(header)
	if err != nil || validApiKey == "" {
		return nil, utils.NewDomainError(401, fmt.Sprintf("api key is required. %s", err))
	}

	if validApiKey != *apiKeyString {
		return nil, utils.NewDomainError(401, fmt.Sprintf("invalid api key. %s", err))
	}

	if event := body.Event; event != "user.upgraded" {
		return nil, utils.NewDomainError(204, "invalid event")
	}

	userID, err := uuid.Parse(body.Data.UserID)
	if err != nil {
		return nil, utils.NewDomainError(401, "user id is required")
	}

	return &User{
		ID:          userID,
		IsChirpyRed: body.Event == "user.upgraded",
	}, nil
}
