package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID   `gorm:"primaryKey;type:uuid;" json:"user_id"`
	Username       string      `gorm:"type:varchar(255)" json:"username"`
	Email          string      `gorm:"type:varchar(255);uniqueIndex;" json:"email"`
	Password       string      `gorm:"type:varchar(255)" json:"password"`
	RegistrationAt time.Time   `gorm:"autoCreateTime" json:"regester_at"`
	LastLogin      time.Time   `json:"last_login"`
	Role           string      `json:"role"`
	Active         bool        `json:"active"`
	EmailVerified  bool        `json:"email_verified"`
	Token          Token       `gorm:"foreignkey:UserID" json:"tokens"`
	Profile        Profile     `gorm:"foreignkey:UserID" json:"profile"`
	SocialLogin    SocialLogin `gorm:"foreignkey:UserID" json:"social_keys"`

	Posts    []Post    `gorm:"foreignkey:UserID"  json:"posts"`
	Comments []Comment `gorm:"foreignkey:UserID" json:"comments"`
	Likes    []Like    `gorm:"foreignkey:UserID" json:"likes"`

	Subscribed bool `gorm:"type:bool;default:false"`
}
type Token struct {
	UserID              uuid.UUID `gorm:"unique,type:char(36);primaryKey" json:"user_id"`
	RefreshToken        string    `json:"refresh_token"`
	ForgotPasswordToken string    `json:"forgot_password_token"`
	VerifyEmailToken    string    `json:"verify_email_token"`
}

type Profile struct {
	UserID  uuid.UUID `gorm:"type:char(36);primaryKey" json:"user_id"`
	Picture string    `json:"picture"`
	Name    string    `gorm:"type:varchar(255)"`
	Bio     string    `json:"bio"`
}

type SocialLogin struct {
	UserID   uuid.UUID `gorm:"type:char(36);foreignKey:ID" json:"user_id"`
	Provder  string    `json:"provider"`
	LoginKey string    `json:"key"`
}

type Post struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;" json:"post_id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	CoverImage  string    `json:"cover_image"`
	PublishDate time.Time
	UserID      uuid.UUID  `gorm:"not null;type:uuid;constraint:OnDelete:CASCADE;"`
	User        User       `gorm:"foreignkey:UserID"`
	Comments    []Comment  `gorm:"foreignkey:PostID;constraint:OnDelete:CASCADE;"`
	Likes       []Like     `gorm:"foreignkey:PostID;constraint:OnDelete:CASCADE;"`
	Tags        []Tag      `gorm:"many2many:post_tags;"`
	Categories  []Category `gorm:"many2many:post_categories;"`
}

type Comment struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Content string    `gorm:"type:text;"`
	PostID  uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnDelete:CASCADE;"`
	Post    Post      `gorm:"foreignkey:PostID"`
	UserID  uuid.UUID `gorm:"not null;type:uuid;index"`
	User    User      `gorm:"foreignkey:UserID"`
}
type Like struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;"`
	PostID uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnDelete:CASCADE;"`
	Post   Post      `gorm:"foreignkey:PostID"`
	UserID uuid.UUID `gorm:"not null;type:uuid;index"`
	User   User      `gorm:"foreignkey:UserID"`
}

type Tag struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name  string    `gorm:"type:varchar(255)"`
	Posts []Post    `gorm:"many2many:post_tags;"`
}

type Category struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name  string    `gorm:"type:varchar(255)"`
	Posts []Post    `gorm:"many2many:post_categories;"`
}

type Subscription struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;"`
	UserID uuid.UUID `gorm:"not null;type:uuid"`
	User   User      `gorm:"foreignkey:UserID"`
}
