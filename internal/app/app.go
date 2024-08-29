package app

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/routes"
	"github.com/uptrace/bun"
)

type App struct {
	http *fiber.App
	db   *bun.DB
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	a.initHttp()

	return a, nil
}

func (a *App) initHttp() {
	a.http = fiber.New()

	routes.SetupRoutes(a.http)
}

func (a *App) Run() {
	a.db = database.ConnectDB()
	if a.db != nil {
		fmt.Println("connected")
	}

	a.http.Listen(":8000")
}
