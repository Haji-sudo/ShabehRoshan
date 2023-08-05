package models

import (
	"time"

	"github.com/google/uuid"
)

// Blog represents the Blog Table
type Blog struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"blog_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	AuthorID    uuid.UUID `gorm:"type:char(36);index" json:"author_id"`
	PublishDate time.Time `gorm:"autoCreateTime" json:"publish_date"`
	Tags        []Tag     `gorm:"many2many:blog_tags;" json:"tags"`
	Comments    []Comment `json:"comments"`
	Likes       []Like    `json:"likes"`
	User        User      `gorm:"foreignKey:AuthorID" json:"user"`
}

// Tag represents the Tag Table
type Tag struct {
	ID    uuid.UUID `gorm:"type:char(36);primaryKey" json:"tag_id"`
	Name  string    `json:"name"`
	Blogs []Blog    `gorm:"many2many:blog_tags;" json:"blogs"`
}

// Comment represents the Comment Table
type Comment struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"comment_id"`
	BlogID      uuid.UUID `gorm:"type:char(36);index" json:"blog_id"`
	UserID      uuid.UUID `gorm:"type:char(36);index" json:"user_id"`
	Content     string    `json:"content"`
	CommentDate time.Time `gorm:"autoCreateTime" json:"comment_date"`
}

// Like represents the Like Table
type Like struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"like_id"`
	BlogID   uuid.UUID `gorm:"type:char(36);index" json:"blog_id"`
	UserID   uuid.UUID `gorm:"type:char(36);index" json:"user_id"`
	LikeDate time.Time `gorm:"autoCreateTime" json:"like_date"`
}

// BlogTag represents the BlogTag Table (Many-to-Many Relationship)
type BlogTag struct {
	BlogID uuid.UUID `gorm:"type:char(36);primaryKey" json:"blog_id"`
	TagID  uuid.UUID `gorm:"type:char(36);primaryKey" json:"tag_id"`
}
