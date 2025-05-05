package repository

import (
	"context"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/crypto"
	"ocserv/pkg/database"
	"time"
)

type TokenRepository struct {
	db *gorm.DB
}

type TokenRepositoryInterface interface {
	CreateToken(ctx context.Context, id uint, uid string, rememberMe bool, isAdmin bool) (string, error)
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		db: database.Get(),
	}
}

func (t *TokenRepository) CreateToken(ctx context.Context, id uint, uid string, rememberMe bool, isAdmin bool) (string, error) {
	expire := time.Now().Add(24 * time.Hour)
	if rememberMe {
		expire = expire.AddDate(0, 1, 0)
	}

	access, err := crypto.GenerateAccessToken(uid, expire.Unix(), isAdmin)
	if err != nil {
		return "", err
	}

	err = t.db.WithContext(ctx).Create(
		&models.UserToken{
			UserID:   id,
			Token:    access,
			ExpireAt: expire,
		},
	).Error
	if err != nil {
		return "", err
	}
	return access, nil
}
