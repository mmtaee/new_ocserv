package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	OcUser       bool `json:"oc_user"`
	OcGroup      bool `json:"oc_group"`
	Statistic    bool `json:"statistic"`
	Occtl        bool `json:"occtl"`
	System       bool `json:"system"`
	SeeServerLog bool `json:"see_server_log"`
}

type User struct {
	ID          uint       `json:"-" gorm:"primaryKey;autoIncrement"`
	UID         string     `json:"uid" gorm:"type:varchar(26);not null;unique"`
	Username    string     `json:"username" gorm:"type:varchar(16);not null;unique"`
	Password    string     `json:"-" gorm:"type:varchar(64); not null"`
	IsAdmin     bool       `json:"is_admin" gorm:"type:bool;default(false)"`
	Salt        string     `json:"-" gorm:"type:varchar(8);not null;unique"`
	LastLogin   *time.Time `json:"last_login"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Token       []UserToken `json:"-"`
	Permissions *Permission `json:"permission" gorm:"type:jsonb"`
}

type UserToken struct {
	ID        uint       `json:"-" gorm:"primaryKey;autoIncrement"`
	UserID    uint       `json:"-" gorm:"index"`
	UID       string     `json:"uid" gorm:"type:varchar(26);not null;unique"`
	Token     string     `json:"token" gorm:"type:varchar(128)"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	ExpireAt  *time.Time `json:"expire_at"`
	User      User       `json:"user"`
}

func (p *Permission) Value() (driver.Value, error) {
	return json.Marshal(&p)
}

func (p *Permission) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &p)
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Permissions == nil {
		u.Permissions = &Permission{
			OcUser:       false,
			OcGroup:      false,
			Statistic:    false,
			Occtl:        false,
			System:       false,
			SeeServerLog: false,
		}
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = ulid.Make().String()
	return
}

func (t *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	t.UID = ulid.Make().String()
	return
}
