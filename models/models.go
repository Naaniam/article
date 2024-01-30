package models

import (
	// built-in pacakge
	"time"

	"gorm.io/gorm"
)

// article details
type Article struct {
	gorm.Model
	ID       uint      `json:"id" gorm:"primaryKey;column:id"`
	Nickname string    `json:"nickname" gorm:"column:nickname" validate:"required"`
	Title    string    `json:"title" gorm:"column:title;not null" validate:"required"`
	Content  string    `json:"content" gorm:"column:content;not null" validate:"required"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:ArticleID"`
}

// comment details
type Comment struct {
	ID        uint    `json:"id" gorm:"primaryKey;column:id"`
	ArticleID uint    `json:"article_id" gorm:"column:article_id"`
	Nickname  string  `json:"nickname" gorm:"column:nickname" validate:"required"`
	Content   string  `json:"content" gorm:"column:content" validate:"required"`
	Replies   []Reply `json:"replies,omitempty" gorm:"foreignKey:CommentID"`
}

// reply comment details
type Reply struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primaryKey column:id"`
	CommentID    uint      `json:"comment_id" gorm:"column:comment_id"`
	Nickname     string    `json:"nickname" gorm:"column:nickname" validate:"required"`
	Content      string    `json:"content" gorm:"column:content" validate:"required"`
	CreationDate time.Time `json:"creation_date" gorm:"column:creation_date"`
}
