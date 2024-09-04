package main

import (
	"github.com/rkashapov2015/webproject/internal/console"
	"github.com/rkashapov2015/webproject/internal/database"
)

func main() {
	database.ConnectDB()
	app := console.New(database.DB)
	app.Run()
}
