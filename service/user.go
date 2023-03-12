package service

import (
	"log"

	"github.com/natepboat/go-rest-api/model"
	"github.com/natepboat/go-rest-api/provider"
)

type IUserService interface {
	GetUser(id int) model.User
}

type UserService struct {
	repository map[int]*model.User
}

func NewUserService(provider *provider.ComponentProvider) *UserService {
	return &UserService{
		repository: provider.Required("UserRepository").(map[int]*model.User),
	}
}

func (s *UserService) GetUser(id int) model.User {
	user := s.repository[id]
	log.Printf("Get user id %v, exist %v \n", id, user != nil)

	return *user
}
