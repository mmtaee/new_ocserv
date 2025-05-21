package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"ocserv/pkg/database"
	"ocserv/pkg/oc"
	"ocserv/pkg/request"
	"slices"
	"strings"
)

type OcservUserRepository struct {
	db          *gorm.DB
	ocUserRepo  oc.OcservUserServiceInterface
	ocOcctlRepo oc.OcctlServiceInterface
}

type OcservUserRepositoryInterface interface {
	GetUsers(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error)
	GetUsersWithOnlineAttr(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error)
	CreateUser(ctx context.Context, user *oc.OcservUser) (*oc.OcservUser, error)
}

func NewOcservUserRepository() *OcservUserRepository {
	return &OcservUserRepository{
		db:          database.Get(),
		ocUserRepo:  oc.NewOcservUserService(database.Get()),
		ocOcctlRepo: oc.NewOcctlService(),
	}
}

// GetUsers get list of users with offset and ordering
func (o *OcservUserRepository) GetUsers(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error) {
	var totalRecords int64
	if err := o.db.WithContext(ctx).Model(&oc.OcservUser{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	order := fmt.Sprintf("%s %s", pagination.Order, pagination.Sort)

	var ocservUsers []oc.OcservUser
	err := o.db.WithContext(ctx).Model(&ocservUsers).Order(order).Limit(pagination.PageSize).Offset(offset).Scan(&ocservUsers).Error
	if err != nil {
		return nil, 0, err
	}
	return &ocservUsers, totalRecords, nil
}

func (o *OcservUserRepository) GetUsersWithOnlineAttr(ctx context.Context, pagination *request.Pagination) (*[]oc.OcservUser, int64, error) {
	type result struct {
		online []string
		err    error
	}

	ch := make(chan result, 1)

	go func() {
		online, err := o.ocOcctlRepo.WithContext(ctx).GetOnlineUsers()
		ch <- result{online: online, err: err}
	}()

	users, count, err := o.GetUsers(ctx, pagination)
	if err != nil {
		return nil, 0, err
	}

	if count > 0 {
		r := <-ch
		if r.err == nil {
			for i, user := range *users {
				if slices.Contains(r.online, user.Username) {
					(*users)[i].IsOnline = true
				}
			}
		}
	}

	return users, count, nil
}

func (o *OcservUserRepository) CreateUser(ctx context.Context, user *oc.OcservUser) (*oc.OcservUser, error) {
	err := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			// create error
			log.Println("create error: ", err)
			return err
		}

		if err := o.ocUserRepo.CreateUser(ctx, user.Username, user.Password, user.Group); err != nil {
			log.Println("CreateUser in os rollback error: ", err)
			// rollback create
			return err
		}

		// commit transaction
		return nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, errors.New("user with that user name already exists")
		}
		log.Println("return transaction error: ", err)
		return nil, err
	}
	return user, nil
}
