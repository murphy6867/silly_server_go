package user

import (
	"encoding/json"
	"net/http"

	"github.com/murphy6867/server/utils"
)

type UserHandler struct {
	svc *UserService
}

func NewUserHandler(svc *UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dataIn CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dataIn); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.svc.CreateUserService(r.Context(), dataIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := ResponseCreateUerDTO{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Email:     user.Email,
	}

	utils.WriteJSON(w, http.StatusCreated, out)
}
