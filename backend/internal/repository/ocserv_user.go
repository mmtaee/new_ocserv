package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	GetUserByID(ctx context.Context, userID string) (*oc.OcservUser, error)
	LockUser(ctx context.Context, userID string, lock bool) error
	UpdateUser(ctx context.Context, ocUser *oc.OcservUser) (*oc.OcservUser, error)
	DeleteUser(ctx context.Context, userID string) error
	DisconnectUser(ctx context.Context, username string) error
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
			return err
		}

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

func (o *OcservUserRepository) GetUserByID(ctx context.Context, userID string) (*oc.OcservUser, error) {
	var ocUser oc.OcservUser
	err := o.db.WithContext(ctx).Where("uid = ?", userID).First(&ocUser).Error
	if err != nil {
		return nil, err
	}
	return &ocUser, nil
}

func (o *OcservUserRepository) LockUser(ctx context.Context, userID string, lock bool) error {
	return o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var ocUser oc.OcservUser

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("uid = ?", userID).
			First(&ocUser).Error; err != nil {
			return err
		}

		if err := o.ocUserRepo.LockUser(ctx, ocUser.Username, lock); err != nil {
			return err
		}

		ocUser.IsLocked = lock
		if err := tx.Save(&ocUser).Error; err != nil {
			return err
		}

		return nil
	})
}

func (o *OcservUserRepository) UpdateUser(ctx context.Context, ocUser *oc.OcservUser) (*oc.OcservUser, error) {
	err := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := o.ocUserRepo.Update(ctx, ocUser.Username, ocUser.Password, ocUser.Group); err != nil {
			return err
		}

		if err := tx.Save(&ocUser).Error; err != nil {
			return err
		}

		return nil
	})
	return ocUser, err
}

func (o *OcservUserRepository) DeleteUser(ctx context.Context, userID string) error {
	err := o.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var ocUser oc.OcservUser

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("uid = ?", userID).
			First(&ocUser).Error; err != nil {
			return err
		}

		if err := tx.Delete(&ocUser).Error; err != nil {
			return err
		}

		if err := o.ocUserRepo.DeleteUser(ctx, ocUser.Username); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (o *OcservUserRepository) DisconnectUser(ctx context.Context, username string) error {
	s, err := o.ocOcctlRepo.WithContext(ctx).DisconnectUser(username)

	log.Println("err: ", err)
	log.Println("result: ", s)

	if err != nil {
		return errors.New("failed to disconnect user " + username)
	}
	return nil
}
