package config

import (
	"log"
	"net/http"

	goappenv "github.com/natepboat/go-app-env"
	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/interceptor"
	"github.com/natepboat/go-rest-api/provider"
	gorouter "github.com/natepboat/go-router"
	"github.com/natepboat/go-router/httpMethod"
)

func ConfigServer(app goappenv.IAppEnv, provider *provider.ComponentProvider) *http.Server {
	r := gorouter.NewRouter(app, nil)
	authInterceptor := provider.Required("interceptor.AuthInterceptor").(interceptor.IHttpInterceptor)
	metricInterceptor := provider.Required("interceptor.MetricInterceptor").(interceptor.IHttpInterceptor)

	s := provider.Required("api.StatController").(*api.StatController)
	u := provider.Required("api.UserController").(*api.UserController)

	r.AddRoute(httpMethod.GET, "/", interceptor.Intercept(s.Home, metricInterceptor))
	r.AddRoute(httpMethod.GET, "/user/:id", u.GetUser)
	r.AddRoute(httpMethod.GET, "/v2/user/:id", interceptor.Intercept(u.GetUser, authInterceptor, metricInterceptor))

	server, err := r.NewServer()
	if err != nil {
		log.Fatalln("Cannot create route server", err)
	}
	return server
}
