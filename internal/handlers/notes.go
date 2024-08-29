package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/database/models"
)

func ListNotes(c *fiber.Ctx) error {
	notes := make([]models.Note, 0)

	if database.DB == nil {
		return SystemErrorMessageJson(c, "Не установлено подключение к БД")
	}

	count, err := database.DB.NewSelect().
		Model(&notes).
		Relation("Author").
		Limit(20).
		ScanAndCount(c.Context())
	if err != nil {
		panic(err)
	}

	return c.Status(200).JSON(
		ResponseRows{
			Count: count,
			Rows:  notes,
		})
}

func ShowNote(c *fiber.Ctx) error {
	note := new(models.Note)
	id := c.Params("id")
	status := fiber.StatusAccepted

	err := database.DB.NewSelect().Model(note).Where("id = ?", id).Scan(c.Context())

	if err != nil {
		return NotFoundMessageJson(c, err.Error())
	}

	return c.Status(status).JSON(note)
}
