package webhook

import (
	"encoding/json"
	"log"
	"net/http"

	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type WebhookHandler struct {
	svc *WebhookService
}

func NewWebhookHandler(svc *WebhookService) *WebhookHandler {
	return &WebhookHandler{svc: svc}
}

func (h *WebhookHandler) UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateSubscriptionDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.HandleError(w, err)
		return
	}

	if err := h.svc.SubscriptionService(r.Context(), r.Header, &req); err != nil {
		log.Println("--------> err", err)

		utils.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
