package config

import (
	"context"
	"log"
	"os"

	goappenv "github.com/natepboat/go-app-env"
)

func ConfigEnv(ctx context.Context) *goappenv.AppEnv {
	app := goappenv.NewAppEnv(os.DirFS("."), ctx)
	log.Println("Application running with active environment = ", app.ActiveEnv())
	log.Println("Application config directory = ", app.ConfigDir())

	return app
}
