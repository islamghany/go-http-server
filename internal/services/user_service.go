package services

import (
	"context"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/models"
)

// UserService defines methods for user operations.
type UserService struct {
	queries *db.Queries
}

// NewUserService creates a new UserService.
func NewUserService(store *db.Queries) *UserService {
	return &UserService{
		queries: store,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*models.User, error) {
	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name:  name,
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
