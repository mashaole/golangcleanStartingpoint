package service

import (
	"errors"
	"goTut/entity"
	"goTut/repository"
	"math/rand"
)

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService(repor repository.PostRepository) PostService {
	repo = repor
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("Title is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}
func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
