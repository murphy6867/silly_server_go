package handler

import (
	"sync/atomic"

	"github.com/murphy6867/silly_server_go/internal/database"
)

type APIConfig struct {
	FileServerHits atomic.Int32
	DB             *database.Queries
	Email          string `json:"email"`
}
