package service

import (
	"github.com/Karibu/api-go-human/model"
	"github.com/Karibu/api-go-human/repository"
)

type IUserService interface {
	repository.IUserRepository
	GetById(string) (model.User, error)
	GetAll() ([]model.User, error)
	Save(model.User) (model.User, error)
}

type UserService struct {
	repository repository.IUserRepository
}

func (s *UserService) Save(m model.User) (model.User, error) {
	return s.repository.Save(m)
}

func (s *UserService) GetById(id string) (model.User, error) {
	return s.repository.GetById(id)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repository.GetAll()
}

func NewUserService(r repository.IUserRepository) (IUserService, error) {
	return &UserService{
		repository: r,
	}, nil
}
