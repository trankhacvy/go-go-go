package posts

import (
	// "errors"
	// "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "os"
)

type PostService struct {
	repo *MySqlPostRepository
}

func NewPostService(r *MySqlPostRepository) *PostService {
	return &PostService{ repo: r }
}

func (service *PostService) Create(post *Post) (bool, error) {
	return service.repo.Create(post)
}

func (service *PostService) Update(post *Post) (bool, error) {
	return service.repo.Update(post)
}

func (service *PostService) Delete(id uint) (bool, error) {
	return service.repo.Delete(id)
}


func (service *PostService) AllPosts() ([]*Post, error) {
	return service.repo.FetchAll()
}