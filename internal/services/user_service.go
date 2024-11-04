package services

import (
	"context"
	db "httpserver/internal/db/sqlc"
	"httpserver/internal/models"
)

// UserService defines methods for user operations.
type UserService struct {
	store *db.Store
}

// NewUserService creates a new UserService.
func NewUserService(store *db.Store) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, email string, hashedPassword []byte) (*models.User, error) {
	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        int(user.ID),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time.UTC().String(),
		UpdatedAt: user.UpdatedAt.Time.UTC().String(),
	}, nil
}
