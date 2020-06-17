package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type VelogConfig struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Title     string     `json:"title"`
	LogoImage string     `json:"logo_image"`
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:time"sql:"index"json:"deleted_at"`
}

func (v VelogConfig) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":         v.ID,
		"title":      v.Title,
		"logo_image": v.LogoImage,
	}
}
