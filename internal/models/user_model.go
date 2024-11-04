package models

type CreateUserParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int    `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
