package main

import (
	"log"

	"github.com/natepboat/go-rest-api/api"
	gorouter "github.com/natepboat/go-router"
	"github.com/natepboat/go-router/httpMethod"
)

func main() {
	r := gorouter.NewRouter(nil, nil)

	r.AddRoute(httpMethod.GET, "/", api.Home)
	r.AddRoute(httpMethod.GET, "/user/:id", api.GetUser)
	server, err := r.NewServer()
	if err != nil {
		log.Fatalln("Cannot create route server", err)
	}
	log.Fatalln(server.ListenAndServe())
}
