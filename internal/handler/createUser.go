package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/murphy6867/server/internal/database"
	"github.com/murphy6867/server/utils"
)

type RegisterRequest struct {
	Email string
}

type UserResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}

func (cfg *APIConfig) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error encoding parameter: %s\n", err)
		utils.WriteJSON(w, 500, map[string]string{"error": "Something went wrong"})
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     req.Email,
	})
	if err != nil {
		log.Printf("Error creating user: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"id":         user.ID.String(),
		"created_at": user.CreatedAt.String(),
		"updated_at": user.UpdatedAt.String(),
		"email":      user.Email,
	})
}
