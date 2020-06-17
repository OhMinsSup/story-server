package models

import (
	"time"
)

type SocialAccount struct {
	ID          string     `gorm:"primary_key;uuid"json:"id"`
	SocialId    string     `sql:"index"json:"social_id"`
	AccessToken string     `json:"access_token"`
	Provider    string     `sql:"index"json:"provider"`
	User        User       `gorm:"foreignkey:UserID"json:"user"`
	UserID      string     `json:"user_id"`
	CreatedAt   time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:time"sql:"index"json:"deleted_at"`
}
