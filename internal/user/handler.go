package user

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// func (h *UserHandler) UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
// 	var req UpdateSubscriptionDTO
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		utils.HandleError(w, err)
// 	}

// 	if err := h.svc.SubscriptionService(r.Context(), &req); err != nil {
// 		utils.HandleError(w, err)
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
