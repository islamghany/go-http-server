package models

type User struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	// Other fields
}
