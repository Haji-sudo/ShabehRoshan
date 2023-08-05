package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the User Table
type User struct {
	ID               uuid.UUID   `gorm:"unique,type:char(36);primaryKey" json:"user_id"`
	Username         string      `json:"username"`
	Email            string      `json:"email"`
	PasswordHash     string      `json:"password_hash"`
	RegistrationDate time.Time   `gorm:"autoCreateTime" json:"registration_date"`
	LastLogin        time.Time   `json:"last_login"`
	Role             string      `json:"role"`
	Active           bool        `json:"active"`
	EmailVerified    bool        `json:"email_verified"`
	Blogs            []Blog      `gorm:"foreignKey:AuthorID" json:"blogs"`
	Comments         []Comment   `gorm:"foreignKey:UserID" json:"comments"`
	Likes            []Like      `gorm:"foreignKey:UserID" json:"likes"`
	Token            Token       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE" json:"token"`
	Profile          Profile     `gorm:"foreignKey:UserID" json:"user_profile"`
	SocialLogin      SocialLogin `gorm:"foreignKey:UserID" json:"user_social_login"`
}

// Token represents user token information
type Token struct {
	UserID              uuid.UUID `gorm:"unique,type:char(36);primaryKey" json:"user_id"`
	RefreshToken        string    `json:"refresh_token"`
	ForgotPasswordToken string    `json:"forgot_password_token"`
	VerifyEmailToken    string    `json:"verify_email_token"`
}

// Profile represents additional user profile information
type Profile struct {
	UserID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"user_id"`
	ProfilePicture string    `json:"profile_picture"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Bio            string    `json:"bio"`
}

// SocialLogin represents social login provider information
type SocialLogin struct {
	UserID            uuid.UUID `gorm:"type:char(36);foreignKey:ID" json:"-"`
	SocialLoginGoogle string    `json:"social_login_google"`
}
