package posts

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type PostRepository interface {
	// FetchAll(int page, string orderBy, string order) ([]*Post, error)
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

func (repo *MySqlPostRepository) FetchAll(page int, limit int, sort string, order string) ([]*Post, int, error) {
	posts := make([]*Post, limit)
	offset := (page - 1) * limit
	count := 0

	db := repo.DB.Offset(offset).Limit(limit)

	if err := db.Find(&posts).Order(fmt.Sprintf("%s %s", sort, order)).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return posts, count, nil
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
