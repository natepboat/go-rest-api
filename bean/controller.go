package bean

import (
	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/service"
)

const (
	StatController    = "api.StatController"
	UserController    = "api.UserController"
	AuthInterceptor   = "interceptor.AuthInterceptor"
	MetricInterceptor = "interceptor.MetricInterceptor"
)

func CreateUserController(beanContext *BeanContext) *api.UserController {
	userService := beanContext.RequiredBean(UserService)
	return api.NewUserController(userService.(service.IUserService))
}
