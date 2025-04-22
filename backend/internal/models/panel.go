package models

import (
	"errors"
	"gorm.io/gorm"
	"ocserv/pkg/database"
)

type Panel struct {
	ID                     uint   `gorm:"primaryKey"`
	Setup                  bool   `json:"setup" gorm:"default:false"`
	GoogleCaptchaSecretKey string `json:"google_captcha_secret" gorm:"type:text"`
	GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" gorm:"type:text"`
}

func (s *Panel) BeforeCreate(tx *gorm.DB) error {
	ch := make(chan error, 1)
	go func() {
		var config Panel
		db := database.Get()
		err := db.Table("panels").First(&config).Error
		if err != nil && config.ID == 0 {
			ch <- nil
		}
		ch <- errors.New("panel configs already exist")
	}()
	return <-ch
}
