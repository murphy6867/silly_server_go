package chirp

import (
	"time"

	"github.com/murphy6867/silly_server_go/internal/database"
)

func mapDatabaseChirps(chs []database.Chirp) ResponseChirpsDTO {
	out := make(ResponseChirpsDTO, len(chs))
	for i, db := range chs {
		out[i] = mapDatabaseChirp(db)
	}
	return out
}

func mapDatabaseChirp(ch database.Chirp) ResponseCreateChirpDTO {
	return ResponseCreateChirpDTO{
		ID:        ch.ID.String(),
		UserID:    ch.UserID.String(),
		Body:      ch.Body,
		CreatedAt: ch.CreatedAt.Format(time.RFC3339),
		UpdatedAt: ch.UpdatedAt.Format(time.RFC3339),
	}
}
