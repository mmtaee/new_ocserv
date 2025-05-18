package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"log"
	"time"
)

type Permission struct {
	OcUsers       bool `json:"oc_users" validate:"omitempty"`
	OcGroups      bool `json:"oc_groups" validate:"omitempty"`
	Statistic     bool `json:"statistic" validate:"omitempty"`
	Occtl         bool `json:"occtl" validate:"omitempty"`
	System        bool `json:"system" validate:"omitempty"`
	SeeServerLogs bool `json:"see_server_logs" validate:"omitempty"`
}

type User struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement" validate:"required"`
	UID         string     `json:"uid" gorm:"type:varchar(26);not null;unique" validate:"required"`
	Username    string     `json:"username" gorm:"type:varchar(16);not null;unique"  validate:"required"`
	Password    string     `json:"-" gorm:"type:varchar(64); not null"`
	IsAdmin     bool       `json:"is_admin" gorm:"type:bool;default(false)"  validate:"required"`
	Salt        string     `json:"-" gorm:"type:varchar(8);not null"`
	LastLogin   *time.Time `json:"last_login"  validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Token       []UserToken `json:"-"`
	Permissions *Permission `json:"permission" gorm:"type:jsonb"  validate:"required"`
}

type UserToken struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"-" gorm:"index"`
	UID       string    `json:"uid" gorm:"type:varchar(26);not null;unique"`
	Token     string    `json:"token" gorm:"type:varchar(128)"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	ExpireAt  time.Time `json:"expire_at"`
	User      User      `json:"user"`
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
	log.Println("u: ", u.IsAdmin)
	if u.IsAdmin {
		u.Permissions = &Permission{
			OcUsers:       true,
			OcGroups:      true,
			Statistic:     true,
			Occtl:         true,
			System:        true,
			SeeServerLogs: true,
		}
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = ulid.Make().String()
	return
}

func (t *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	if t.UID == "" {
		t.UID = ulid.Make().String()
	}
	return
}
