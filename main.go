package main

import (
	"context"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/natepboat/go-rest-api/bean"
	"github.com/natepboat/go-rest-api/config"
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
	appContext := context.Background()
	app := config.ConfigEnv(appContext)

	beanContext := bean.InitBeanContext()
	server := config.ConfigServer(app, beanContext)

	log.Fatalln(server.ListenAndServe())
}
