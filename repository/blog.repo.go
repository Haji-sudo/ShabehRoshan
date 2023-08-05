package repository

import (
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetByID(id uuid.UUID) (*models.Blog, error)
	GetByTitle(title string) ([]models.Blog, error)
	Create(blog *models.Blog) error
	Update(blog *models.Blog) error
	Delete(blog *models.Blog) error
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

func (r *blogRepo) GetByID(id uuid.UUID) (*models.Blog, error) {
	blog := new(models.Blog)

	return blog, nil
}

func (r *blogRepo) GetByTitle(title string) ([]models.Blog, error) {
	blogs := []models.Blog{}
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&blogs).Error
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
func (r *blogRepo) Create(blog *models.Blog) error {

	return r.db.Create(blog).Error
}

func (r *blogRepo) Update(blog *models.Blog) error {

	return r.db.Save(blog).Error
}

func (r *blogRepo) Delete(blog *models.Blog) error {

	return r.db.Delete(blog).Error
}
