package service

import (
	"log"

	"github.com/natepboat/go-rest-api/model"
)

type IUserService interface {
	GetUser(id int) model.User
}

type UserService struct {
	repository map[int]*model.User
}

func NewUserService(repository map[int]*model.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUser(id int) model.User {
	user := s.repository[id]
	log.Printf("Get user id %v, exist %v \n", id, user != nil)

	return *user
}
