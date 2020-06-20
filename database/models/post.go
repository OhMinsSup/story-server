package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type Post struct {
	ID         string    `gorm:"primary_key;uuid"json:"id"`
	Title      string    `json:"title"`
	Body       string    `gorm:"type:text"json:"body"`
	Thumbnail  string    `json:"thumbnail"`
	IsMarkdown bool      `json:"is_markdown"`
	IsTemp     bool      `json:"is_temp"`
	IsPrivate  bool      `gorm:"default:true"json:"is_private"`
	Likes      int       `gorm:"default:0"json:"likes"`
	Views      int       `gorm:"default:0"json:"views"`
	User       User      `gorm:"foreignkey:UserID"json:"user"`
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	//DeletedAt   *time.Time    `sql:"index"json:"deleted_at"`
	Tags        []Tag         `gorm:"many2many:posts_tags;association_foreignkey:tag_id;foreignkey:post_id;"json:"tags"`
	PostScore   []PostScore   `gorm:"polymorphic:Owner;"`
	PostRead    []PostRead    `gorm:"polymorphic:Owner;"`
	PostHistory []PostHistory `gorm:"polymorphic:Owner;"`
	PostLike    []PostLike    `gorm:"polymorphic:Owner;"`
}

// Serialize serializes post data
func (p Post) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":          p.ID,
		"user_id":     p.UserID,
		"title":       p.Title,
		"body":        p.Body,
		"thumbnail":   p.Thumbnail,
		"likes":       p.Likes,
		"views":       p.Views,
		"is_markdown": p.IsMarkdown,
		"is_temp":     p.IsTemp,
		"is_private":  p.IsPrivate,
		"created_at":  p.CreatedAt,
		"updated_at":  p.UpdatedAt,
	}
}
