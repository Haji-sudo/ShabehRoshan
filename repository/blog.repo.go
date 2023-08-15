package repository

import (
	"log"

	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetByID(id uuid.UUID) (*models.Post, error)
	GetByTitle(title string) ([]models.Post, error)
	Create(blog *models.Post) error
	Update(blog *models.Post) error
	Delete(blog *models.Post) error
	Get10LastPost() []models.Post
}

type blogRepo struct {
	db *gorm.DB
}

func NewBlogRepository() BlogRepository {
	repo := &blogRepo{
		db: db.DB,
	}
	return repo
}

func (r *blogRepo) GetByID(id uuid.UUID) (*models.Post, error) {
	blog := new(models.Post)
	return blog, nil
}

func (r *blogRepo) GetByTitle(title string) ([]models.Post, error) {
	blogs := []models.Post{}
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
func (r *blogRepo) Create(blog *models.Post) error {

	return r.db.Create(blog).Error
}

func (r *blogRepo) Update(blog *models.Post) error {

	return r.db.Save(blog).Error
}

func (r *blogRepo) Delete(blog *models.Post) error {

	return r.db.Delete(blog).Error
}

func (r *blogRepo) Get10LastPost() []models.Post {
	var posts []models.Post
	result := r.db.Order("publish_date desc").Limit(10).Find(&posts)
	if result.Error != nil {
		log.Printf("Error fetching records: %v", result.Error)
	}
	return posts
}
