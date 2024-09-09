package main

import (
	"context"
	"log"

	"github.com/rkashapov2015/webproject/internal/app"
	"github.com/rkashapov2015/webproject/internal/config"
)

func main() {
	ctx := context.Background()

	config.LoadEnv(".env")
	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Run()
}
