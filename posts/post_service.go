package posts

import (
	
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

func (service *PostService) AllPosts(page int, limit int, sort string, order string) ([]*Post, int, error) {
	return service.repo.FetchAll(page, limit, sort, order)
}