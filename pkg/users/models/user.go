//go:generate mockgen -source=user.go -destination=../../../internal/mocks/users/repository/user.go -package=mock_users_repository UserRepository
package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"uniqueIndex"`
	Password string
	Email    string `gorm:"uniqueIndex"`
	Name     string
}

// UserRepository defines the data access operations for users.
type UserRepository interface {
	// Create creates a new user.
	Create(context.Context, *User) error
	// GetByUsername returns a user by username.
	GetByUsername(context.Context, string) (*User, error)
	// GetByEmail returns a user by email.
	GetByEmail(context.Context, string) (*User, error)
	// GetByID returns a user by ID.
	GetByID(context.Context, string) (*User, error)
	// Update updates a user.
	Update(context.Context, *User) error
	// Delete deletes a user.
	Delete(context.Context, *User) error
}
