package posts

import (
	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	FetchAll() ([]*Post, error)
	Create(*Post) (bool, error)
	Update(*Post) (bool, error)
	Delete(id uint) (bool, error)
	FindById(id uint) (*Post, error)
}

type MySqlPostRepository struct {
	DB *gorm.DB
}

func NewMySqlPostRepository(db *gorm.DB) *MySqlPostRepository {
	return &MySqlPostRepository{DB: db}
}

func (repo *MySqlPostRepository) FetchAll() ([]*Post, error) {
	posts := make([]*Post, 0)
	if err := repo.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (repo *MySqlPostRepository) Create(post *Post) (bool, error) {
	if err := repo.DB.Create(post).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *MySqlPostRepository) Update(post *Post) (bool, error) {
	if err := repo.DB.Save(post).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *MySqlPostRepository) Delete(id uint) (bool, error) {
	post := &Post{}
	if err := repo.DB.First(post, id).Error; err != nil {
		return false, err
	}
	if err := repo.DB.Delete(post).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repo *MySqlPostRepository) FindById(id uint) (*Post, error) {
	post := &Post{}
	if err := repo.DB.First(post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}
