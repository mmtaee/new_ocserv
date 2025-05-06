package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/crypto"
	"ocserv/pkg/database"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	CreateAdmin(ctx context.Context, user *models.User) (*models.User, error)
	CreateStaff(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	ChangePassword(ctx context.Context, id string, currentPassword, newPassword string) error
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

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.WithContext(ctx).First(&user, "uid = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ChangePassword(ctx context.Context, id string, currentPassword, newPassword string) error {
	user := &models.User{}
	err := r.db.WithContext(ctx).First(&user, "uid = ?", id).Error
	if err != nil {
		return err
	}

	if !crypto.CheckPassword(currentPassword, user.Password, user.Salt) {
		return errors.New("invalid old password")
	}

	passwd := crypto.CreatePassword(newPassword)
	user.Password = passwd.Hash
	user.Salt = passwd.Salt

	return r.db.WithContext(ctx).Save(&user).Error
}
