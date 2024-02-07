package repository

import (
	"context"

	"cbsr.io/golang-grpc-template/modules/users/models"
	"gorm.io/gorm"
)

var _ models.IUserRepository = &userRepository{}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) models.IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return tx.Create(user).Error
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	tx := r.db.WithContext(ctx)

	err := tx.First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return tx.Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, user *models.User) error {
	tx := r.db.WithContext(ctx)
	return tx.Delete(user).Error
}
