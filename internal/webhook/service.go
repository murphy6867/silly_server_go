package webhook

import (
	"context"
	"net/http"

	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type WebhookService struct {
	repo WebHookRepository
}

func NewWebhookService(r WebHookRepository) *WebhookService {
	return &WebhookService{repo: r}
}

func (svc *WebhookService) SubscriptionService(ctx context.Context, header http.Header, body *UpdateSubscriptionDTO) error {
	if event := body.Event; event == "" {
		return utils.NewDomainError(204, "event is required")
	}

	if userID := body.Data.UserID; userID == "" {
		return utils.NewDomainError(204, "user id is required")
	}

	apiKey := svc.repo.GetAPIKeyString()

	updateStatus, err := MappingUpdateSubscription(header, body, apiKey)
	if err != nil {
		return err
	}

	return svc.repo.UpdateSubscriptionChirpyRed(ctx, updateStatus)
}
