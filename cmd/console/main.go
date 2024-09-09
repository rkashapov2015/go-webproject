package main

import (
	"github.com/rkashapov2015/webproject/internal/config"
	"github.com/rkashapov2015/webproject/internal/console"
	"github.com/rkashapov2015/webproject/internal/database"
)

func main() {
	config.LoadEnv(".env")
	database.ConnectDB()
	app := console.New(database.DB)
	app.Run()
}
