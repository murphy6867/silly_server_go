package auth

import (
	"encoding/json"
	"net/http"

	utils "github.com/murphy6867/silly_server_go/internal/shared"
)

type AuthHandler struct {
	svc *AuthService
}

func NewAuthHandler(svc *AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var dataIn SignUpUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dataIn); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.svc.SignUpUserService(r.Context(), dataIn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out := ResponseUerDTO{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Email:     user.Email,
	}

	utils.WriteJSON(w, http.StatusCreated, out)
}

func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInInfo SignInDTO
	if err := json.NewDecoder(r.Body).Decode(&signInInfo); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	out, err := h.svc.SignInService(r.Context(), signInInfo)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ResponseUerDTO{
		ID:        out.ID.String(),
		CreatedAt: out.CreatedAt.String(),
		UpdatedAt: out.UpdatedAt.String(),
		Email:     out.Email,
	})
}
