package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	if database.DB != nil {
		fmt.Println("Setup Routes db not empty")
	}
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)
	app.Patch("/fact/:id", handlers.UpdateFact)

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	notes := v1.Group("/notes")
	notes.Get("/", handlers.ListNotes)
	notes.Post("/", handlers.CreateNote)
	notes.Get("/:id", handlers.ShowNote)
	notes.Patch("/:id", handlers.UpdateNote)
	notes.Delete("/:id", nil)
}
