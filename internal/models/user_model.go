package models

type User struct {
	ID       int64  `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	// Other fields
}
