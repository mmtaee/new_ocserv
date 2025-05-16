package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/crypto"
	"ocserv/pkg/database"
	"ocserv/pkg/request"
)

type UserRepository struct {
	db         *gorm.DB
	cryptoRepo crypto.CustomPasswordInterface
}

type UserRepositoryInterface interface {
	CreateAdmin(ctx context.Context, user *models.User) (*models.User, error)
	CreateStaff(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserById(ctx context.Context, uid string) (*models.User, error)
	ChangePassword(ctx context.Context, uid string, currentPassword, newPassword string) error
	GetStaffs(ctx context.Context, pagination *request.Pagination, filters ...map[string]interface{}) (*[]models.User, int64, error)
	UpdatePermission(ctx context.Context, id uint, perm models.Permission) error
	Delete(ctx context.Context, id uint) error
	ChangeStaffPassword(ctx context.Context, id uint, passwordHash, salt string) error
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db:         database.Get(),
		cryptoRepo: crypto.NewCustomPassword(),
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

func (r *UserRepository) GetUserById(ctx context.Context, uid string) (*models.User, error) {
	user := &models.User{}
	err := r.db.WithContext(ctx).First(&user, "uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) ChangePassword(ctx context.Context, uid string, currentPassword, newPassword string) error {
	user := &models.User{}
	err := r.db.WithContext(ctx).First(&user, "uid = ?", uid).Error
	if err != nil {
		return err
	}

	if !r.cryptoRepo.CheckPassword(currentPassword, user.Password, user.Salt) {
		return errors.New("invalid old password")
	}

	passwd := r.cryptoRepo.CreatePassword(newPassword)
	user.Password = passwd.Hash
	user.Salt = passwd.Salt

	return r.db.WithContext(ctx).Save(&user).Error
}

func (r *UserRepository) GetStaffs(ctx context.Context, pagination *request.Pagination, filters ...map[string]interface{}) (*[]models.User, int64, error) {
	var totalRecords int64

	whereFilters := "is_admin = false"

	if err := r.db.WithContext(ctx).Model(&models.User{}).Where(whereFilters).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if pagination.Order == "" {
		pagination.Order = "id"
		pagination.Sort = "ASC"
	}

	var staffs []models.User
	offset := (pagination.Page - 1) * pagination.PageSize
	order := fmt.Sprintf("%s %s", pagination.Order, pagination.Sort)

	err := r.db.WithContext(ctx).Model(&staffs).Where(whereFilters).Order(order).Limit(pagination.PageSize).Offset(offset).Scan(&staffs).Error
	if err != nil {
		return nil, 0, err
	}
	return &staffs, totalRecords, nil
}

func (r *UserRepository) UpdatePermission(ctx context.Context, id uint, perm models.Permission) error {
	jsonPerm, err := json.Marshal(perm)
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Update("permissions", jsonPerm).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) ChangeStaffPassword(ctx context.Context, id uint, passwordHash, salt string) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"password": passwordHash,
		"salt":     salt,
	}).Error
}
