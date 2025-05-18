package oc

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"ocserv/pkg/request"
)

type OcservUserService struct {
	db *gorm.DB
}

type OcservUserServiceInterface interface {
	GetUsers(ctx context.Context, pagination *request.Pagination) (*[]OcservUser, int64, error)
}

func NewOcservUserService(db *gorm.DB) *OcservUserService {
	return &OcservUserService{db: db}
}

// GetUsers get list of users with offset and ordering
func (o *OcservUserService) GetUsers(ctx context.Context, pagination *request.Pagination) (*[]OcservUser, int64, error) {
	var totalRecords int64
	if err := o.db.WithContext(ctx).Model(&OcservUser{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	order := fmt.Sprintf("%s %s", pagination.Order, pagination.Sort)

	var ocservUsers []OcservUser
	err := o.db.WithContext(ctx).Model(&ocservUsers).Order(order).Limit(pagination.PageSize).Offset(offset).Scan(&ocservUsers).Error
	if err != nil {
		return nil, 0, err
	}
	return &ocservUsers, totalRecords, nil
}
