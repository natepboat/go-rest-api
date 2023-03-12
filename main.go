package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"

	goappenv "github.com/natepboat/go-app-env"
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

func startServer(server *http.Server, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		log.Printf("Application server listen on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Cannot start server, error: %v", err)
		}
	}()
}

func awaitInterupt(server *http.Server, context context.Context) {
	signalChan := make(chan os.Signal)

	go func() {
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		interuptSig := <-signalChan

		log.Printf("Got interupt signal, type: %s, shutting down server...\n", interuptSig)
		if err := server.Shutdown(context); err != nil {
			panic(err)
		}
	}()
}

func main() {
	log.Println("Preparing application context and resources")
	appContext := context.Background()
	app := goappenv.NewAppEnv(os.DirFS("."), appContext)
	provider := config.InitComponentProvider()

	log.Println("Starting application server...")
	serverWg := &sync.WaitGroup{}
	serverWg.Add(1)
	server := config.ConfigServer(app, provider)
	startServer(server, serverWg)
	awaitInterupt(server, appContext)

	serverWg.Wait()
	log.Println("Application gracefully exit")
}
