package repository

import (
	"context"
	"gorm.io/gorm"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

type PanelRepository struct {
	db *gorm.DB
}

type PanelRepositoryInterface interface {
	Create(ctx context.Context, panel *models.Panel) (*models.Panel, error)
	GetConfig(ctx context.Context) (*models.Panel, error)
	Update(ctx context.Context, panel *models.Panel) (*models.Panel, error)
}

func NewPanelRepository() *PanelRepository {
	return &PanelRepository{db: database.Get()}
}

func (p *PanelRepository) GetConfig(ctx context.Context) (*models.Panel, error) {
	panel := &models.Panel{}
	if err := p.db.WithContext(ctx).First(panel).Error; err != nil {
		return nil, err
	}
	return panel, nil
}

func (p *PanelRepository) Create(ctx context.Context, panel *models.Panel) (*models.Panel, error) {
	err := p.db.WithContext(ctx).Create(&panel).Error
	if err != nil {
		return nil, err
	}
	return panel, nil
}

func (p *PanelRepository) Update(ctx context.Context, panel *models.Panel) (*models.Panel, error) {
	var existingPanel models.Panel
	if err := p.db.WithContext(ctx).First(&existingPanel).Error; err != nil {
		return nil, err
	}
	panel.ID = existingPanel.ID
	err := p.db.WithContext(ctx).Save(&panel).Error
	if err != nil {
		return nil, err
	}
	return panel, nil
}
