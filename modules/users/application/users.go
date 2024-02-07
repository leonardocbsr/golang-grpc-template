package application

import (
	"context"

	"cbsr.io/golang-grpc-template/modules/users/models"
)

type IUserService interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, username, password, email, name string) (*models.User, error)
	// GetUserByUsername returns a user by username
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	// GetUserByEmail returns a user by email
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUserByID returns a user by ID
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, user *models.User) error
}

type userService struct {
	repo models.IUserRepository
}

// NewUserService creates a new user service
func NewUserService(repo models.IUserRepository) IUserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, username, password, email, name string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
		Email:    email,
		Name:     name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, user *models.User) error {
	return s.repo.Delete(ctx, user)
}
