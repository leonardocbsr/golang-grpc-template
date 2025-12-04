// Package application provides business logic for user operations.
//
//go:generate mockgen -source=users.go -destination=../../../internal/mocks/users/application/users.go -package=mock_users_application
package application

import (
	"context"

	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService defines the business logic for user operations.
type UserService interface {
	// CreateUser creates a new user with hashed password.
	CreateUser(ctx context.Context, username, password, email, name string) (*models.User, error)
	// VerifyPassword checks if password matches the user's stored hash.
	VerifyPassword(ctx context.Context, username, password string) (*models.User, error)
	// GetUserByUsername returns a user by username.
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	// GetUserByEmail returns a user by email.
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUserByID returns a user by ID.
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	// UpdateUser updates a user.
	UpdateUser(ctx context.Context, user *models.User) error
	// DeleteUser deletes a user.
	DeleteUser(ctx context.Context, user *models.User) error
}

type userService struct {
	repo models.UserRepository
}

// NewUserService creates a new user service.
func NewUserService(repo models.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, username, password, email, name string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Name:     name,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) VerifyPassword(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
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
