package chirp

type CreateChirpDTO struct {
	UserID string `json:"user_id"`
	Body   string `json:"body"`
}

type ResponseCreateChirpDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
}
