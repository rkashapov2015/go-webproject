package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/database/models"
)

func ListNotes(c *fiber.Ctx) error {
	notes := make([]models.Note, 0)

	if database.DB == nil {
		fmt.Println(database.DB)
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

func CreateNote(c *fiber.Ctx) error {
	note := new(models.Note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.NewInsert().Model(note).Exec(c.Context())

	return c.Status(200).JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	note := new(models.Note)
	id := c.Params("id")

	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	_, err := database.DB.NewUpdate().Model(note).Where("id = ?", id).Exec(c.Context())

	if err != nil {
		return c.Status(200).JSON(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(note)
}
