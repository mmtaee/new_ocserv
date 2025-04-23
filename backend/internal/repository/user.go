package repository

import (
	"context"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateAdmin(ctx context.Context, user *models.User) (*models.User, error)
	CreateStaff(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.Get(),
	}
}

func (r *UserRepository) CreateAdmin(ctx context.Context, user *models.User) (*models.User, error) {
	user.IsAdmin = true
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateStaff(ctx context.Context, user *models.User) (*models.User, error) {
	user.IsAdmin = false
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
