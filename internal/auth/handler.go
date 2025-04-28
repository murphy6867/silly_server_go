package auth

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusUnauthorized)
		return
	}

	out := ResponseUerDTO{
		ID:          user.ID.String(),
		CreatedAt:   user.CreatedAt.String(),
		UpdatedAt:   user.UpdatedAt.String(),
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
	}

	utils.WriteJSON(w, http.StatusCreated, out)
}

func (h *AuthHandler) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInInfo SignInDTO
	if err := json.NewDecoder(r.Body).Decode(&signInInfo); err != nil {
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusBadRequest)
		return
	}

	out, err := h.svc.SignInService(r.Context(), signInInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Authorization", "Bearer "+out.Token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := SignInResponse{
		User: User{
			ID:          out.User.ID,
			CreatedAt:   out.User.CreatedAt,
			UpdatedAt:   out.User.UpdatedAt,
			Email:       out.User.Email,
			IsChirpyRed: out.User.IsChirpyRed,
		},
		Token:        out.Token,
		RefreshToken: out.RefreshToken,
		IsChirpyRed:  out.User.IsChirpyRed,
		Email:        out.User.Email,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	out, err := h.svc.RefreshTokenService(r.Context(), r.Header)
	if err != nil {
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusUnauthorized)
		return
	}
	utils.WriteJSON(w, http.StatusOK, RefreshResponse{
		Token: out.RefreshToken,
	})
}

func (h *AuthHandler) RevokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.RevokeRefreshTokenService(r.Context(), r.Header); err != nil {
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *AuthHandler) UpdateEmailAndPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var body EditEmailAndPasswordDTO
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, fmt.Sprintf("no body request: %v", err), http.StatusBadRequest)
		return
	}

	out, err := h.svc.UpdateEmailAndPasswordService(r.Context(), r.Header, body)
	if err != nil {
		http.Error(w, fmt.Sprintf("unauthorized: %v", err), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Authorization", "Bearer "+out.Token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	utils.WriteJSON(w, http.StatusOK, User{
		ID:    out.User.ID,
		Email: out.User.Email,
	})
}
