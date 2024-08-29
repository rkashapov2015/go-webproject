package main

import (
	"context"
	"log"

	"github.com/rkashapov2015/webproject/internal/app"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	app.Run()

	// db := database.ConnectDB()
	// if db != nil {
	// 	fmt.Println("Connected")
	// }
	// fmt.Println(database.DB)
	// app := fiber.New()

	// setupRoutes(app)

	// app.Listen(":8000")
}
