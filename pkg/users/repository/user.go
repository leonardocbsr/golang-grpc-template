// Package repository provides database access for user operations.
package repository

import (
	"context"

	apperrors "github.com/leonardocbsr/golang-grpc-template/internal/errors"
	"github.com/leonardocbsr/golang-grpc-template/pkg/users/models"
	"gorm.io/gorm"
)

var _ models.UserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return apperrors.WrapDBError(tx.Create(user).Error)
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.Where("username = ?", username).First(&user).Error
	return &user, apperrors.WrapDBError(err)
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.Where("email = ?", email).First(&user).Error
	return &user, apperrors.WrapDBError(err)
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.First(&user, id).Error
	return &user, apperrors.WrapDBError(err)
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return apperrors.WrapDBError(tx.Save(user).Error)
}

func (r *userRepository) Delete(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return apperrors.WrapDBError(tx.Delete(user).Error)
}
