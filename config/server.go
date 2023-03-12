package config

import (
	"log"
	"net/http"

	goappenv "github.com/natepboat/go-app-env"
	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/bean"
	"github.com/natepboat/go-rest-api/interceptor"
	gorouter "github.com/natepboat/go-router"
	"github.com/natepboat/go-router/httpMethod"
)

func ConfigServer(app goappenv.IAppEnv, beanContext *bean.BeanContext) *http.Server {
	r := gorouter.NewRouter(app, nil)
	authInterceptor := beanContext.RequiredBean(bean.AuthInterceptor).(interceptor.IHttpInterceptor)
	metricInterceptor := beanContext.RequiredBean(bean.MetricInterceptor).(interceptor.IHttpInterceptor)

	s := beanContext.RequiredBean(bean.StatController).(*api.StatController)
	u := beanContext.RequiredBean(bean.UserController).(*api.UserController)

	r.AddRoute(httpMethod.GET, "/", interceptor.Intercept(s.Home, metricInterceptor))
	r.AddRoute(httpMethod.GET, "/user/:id", u.GetUser)
	r.AddRoute(httpMethod.GET, "/v2/user/:id", interceptor.Intercept(u.GetUser, authInterceptor, metricInterceptor))

	server, err := r.NewServer()
	if err != nil {
		log.Fatalln("Cannot create route server", err)
	}
	return server
}
