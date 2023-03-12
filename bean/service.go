package bean

import (
	"time"

	"github.com/natepboat/go-rest-api/model"
	"github.com/natepboat/go-rest-api/service"
)

const (
	UserService = "service.UserService"
)

func CreateUserService() service.IUserService {
	service := service.NewUserService(map[int]*model.User{
		1: {
			Id:   1,
			Name: "Boat",
			BirthDate: func() time.Time {
				t, _ := time.Parse("yyyy-MM-dd", "2023-01-31")
				return t
			}(),
		},
		2: {
			Id:   2,
			Name: "Za",
			BirthDate: func() time.Time {
				t, _ := time.Parse("yyyy-MM-dd", "2022-02-27")
				return t
			}(),
		},
	})

	return service
}
