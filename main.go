package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/interceptor"
	gorouter "github.com/natepboat/go-router"
	"github.com/natepboat/go-router/httpMethod"
)

func init() {
	go func() {
		time.Sleep(120 * time.Second)
		memprof, err := os.Create("mem.pprof")
		if err != nil {
			log.Fatalln(err)
		}
		pprof.WriteHeapProfile(memprof)
		memprof.Close()
	}()
}

func main() {
	authInterceptor := interceptor.NewAuthInterceptor()
	metricInterceptor := interceptor.NewMetricInterceptor()

	r := gorouter.NewRouter(nil, nil)

	r.AddRoute(httpMethod.GET, "/", interceptor.Intercept(api.Home, metricInterceptor))
	r.AddRoute(httpMethod.GET, "/user/:id", api.GetUser)
	r.AddRoute(httpMethod.GET, "/v2/user/:id", interceptor.Intercept(api.GetUser, authInterceptor, metricInterceptor))
	server, err := r.NewServer()
	if err != nil {
		log.Fatalln("Cannot create route server", err)
	}
	log.Fatalln(server.ListenAndServe())
}
