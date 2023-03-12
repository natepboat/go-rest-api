package config

import (
	"time"

	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/interceptor"
	"github.com/natepboat/go-rest-api/model"
	"github.com/natepboat/go-rest-api/provider"
	"github.com/natepboat/go-rest-api/service"
)

func InitComponentProvider() *provider.ComponentProvider {
	provider := provider.NewComponentProvider()

	provider.Add("UserRepository", NewUserRepository())
	provider.Add("service.UserService", service.NewUserService(provider))
	provider.Add("interceptor.AuthInterceptor", interceptor.NewAuthInterceptor())
	provider.Add("interceptor.MetricInterceptor", interceptor.NewMetricInterceptor())
	provider.Add("api.StatController", api.NewStatController())
	provider.Add("api.UserController", api.NewUserController(provider))

	return provider
}

func NewUserRepository() map[int]*model.User {
	return map[int]*model.User{
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
	}
}
