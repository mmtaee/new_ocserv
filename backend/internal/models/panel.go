package models

import (
	"errors"
	"gorm.io/gorm"
)

type Panel struct {
	ID                     uint   `gorm:"primaryKey"`
	Setup                  bool   `json:"setup" gorm:"default:false"`
	GoogleCaptchaSecretKey string `json:"google_captcha_secret" gorm:"type:text"`
	GoogleCaptchaSiteKey   string `json:"google_captcha_site_key" gorm:"type:text"`
}

func (s *Panel) BeforeCreate(tx *gorm.DB) error {
	var existing Panel

	if err := tx.First(&existing).Error; err == nil {
		return errors.New("a site configuration already exists; only one is allowed")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return err
	}
}
