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
	Active         bool        `json:"active"`
	EmailVerified  bool        `json:"email_verified"`
	Token          Token       `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tokens"`
	Profile        Profile     `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
	SocialLogin    SocialLogin `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_keys"`

	Posts    []Post    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"  json:"posts"`
	Comments []Comment `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Likes    []Like    `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"likes"`

	Subscribed bool `gorm:"type:bool;default:false"`
}
type Token struct {
	UserID              uuid.UUID `gorm:"not null;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RefreshToken        string    `json:"refresh_token"`
	ForgotPasswordToken string    `json:"forgot_password_token"`
	VerifyEmailToken    string    `json:"verify_email_token"`
}

type Profile struct {
	UserID  uuid.UUID `gorm:"not null;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Picture string    `json:"picture"`
	Name    string    `gorm:"type:varchar(255)"`
	Bio     string    `json:"bio"`
}

type SocialLogin struct {
	UserID   uuid.UUID `gorm:"not null;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Provder  string    `json:"provider"`
	LoginKey string    `json:"key"`
}

type Post struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;" json:"post_id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	PublishDate time.Time
	UserID      uuid.UUID  `gorm:"not null;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User        User       `gorm:"foreignkey:UserID"`
	Comments    []Comment  `gorm:"foreignkey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes       []Like     `gorm:"foreignkey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags        []Tag      `gorm:"many2many:post_tags;"`
	Categories  []Category `gorm:"many2many:post_categories;"`
}

type Comment struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Content string    `gorm:"type:text;"`
	PostID  uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Post    Post      `gorm:"foreignkey:PostID"`
	UserID  uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User    User      `gorm:"foreignkey:UserID"`
}
type Like struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid;"`
	PostID uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Post   Post      `gorm:"foreignkey:PostID"`
	UserID uuid.UUID `gorm:"not null;type:uuid;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
	UserID uuid.UUID `gorm:"not null;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User   User      `gorm:"foreignkey:UserID"`
}
