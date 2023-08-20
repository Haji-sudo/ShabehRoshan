package repository

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	m "github.com/haji-sudo/ShabehRoshan/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uuid.UUID) (*m.User, error)
	GetByUsername(username string) (*m.User, error)
	GetByEmail(email string) (*m.User, error)
	Create(user *m.User) error
	Update(user *m.User) error
	Delete(user *m.User) error
	EmailExist(email string) (bool, error)
	UsernameExist(username string) (bool, error)
	GetToken(user *m.User) error
	UpdateToken(user *m.User) error
	GetProfile(user *m.User) error
	UpdateProfile(user *m.User) error

	BatchUpdate(user *m.User) error
	GetAllPosts(userid uuid.UUID) ([]m.Post, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	repo := &userRepo{
		db: db.DB,
	}

	return repo
}

func (r *userRepo) GetByID(id uuid.UUID) (*m.User, error) {
	user := new(m.User)
	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepo) GetByUsername(username string) (*m.User, error) {
	user := new(m.User)
	err := r.db.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepo) GetByEmail(email string) (*m.User, error) {
	user := new(m.User)
	err := r.db.Where("email = ?", strings.ToLower(email)).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) EmailExist(email string) (bool, error) {
	_, err := r.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *userRepo) UsernameExist(username string) (bool, error) {
	_, err := r.GetByUsername(strings.ToLower(username))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *userRepo) GetToken(user *m.User) error {
	err := r.db.Preload("Token").First(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepo) GetProfile(user *m.User) error {
	err := r.db.Preload("Profile").First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) GetAllPosts(userid uuid.UUID) ([]m.Post, error) {
	posts := []m.Post{}
	err := db.DB.Where("author_id = ?", userid).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *userRepo) Create(user *m.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) Update(user *m.User) error {
	return r.db.Save(user).Error
}
func (r *userRepo) BatchUpdate(user *m.User) error {
	return r.db.Save(user).Save(user.Token).Error
}
func (r *userRepo) UpdateToken(user *m.User) error {
	return r.db.Save(user.Token).Error
}

func (r *userRepo) Delete(user *m.User) error {
	return r.db.Delete(user).Error
}
func (r *userRepo) UpdateProfile(user *m.User) error {
	return r.db.Save(user.Profile).Error
}
