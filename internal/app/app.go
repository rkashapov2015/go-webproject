package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/routes"
)

type App struct {
	http *fiber.App
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.initHttp()

	return a, nil
}

func (a *App) initHttp() {
	a.http = fiber.New()
	database.ConnectDB()

	routes.SetupRoutes(a.http)
}

func (a *App) Run() {
	if database.DB == nil {
		panic("DB not connected")
	}
	a.http.Listen(":8000")
}
