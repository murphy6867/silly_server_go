package handler

import (
	"sync/atomic"

	"github.com/murphy6867/server/internal/database"
)

type APIConfig struct {
	FileServerHits atomic.Int32
	DB             *database.Queries
}

type ValidateChirps struct {
	Body string `json:"body"`
}
