package user

type CreateUserDTO struct {
	Email string `json:"email"`
}

type ResponseCreateUerDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Email     string `json:"email"`
}
