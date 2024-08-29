package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rkashapov2015/webproject/internal/database"
	"github.com/rkashapov2015/webproject/internal/database/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home Route")
}

func ListFacts(c *fiber.Ctx) error {
	facts := make([]models.Fact, 0)

	count, err := database.DB.NewSelect().Model(&facts).Limit(20).ScanAndCount(c.Context())
	if err != nil {
		panic(err)
	}

	response := ResponseRows{
		Count: count,
		Rows:  facts,
	}

	return c.Status(200).JSON(response)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.NewInsert().Model(fact).Exec(c.Context())

	return c.Status(200).JSON(fact)
}

func ShowFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")
	status := fiber.StatusAccepted

	err := database.DB.NewSelect().Model(fact).Where("id = ?", id).Scan(c.Context())

	if err != nil {
		return NotFoundMessageJson(c, err.Error())
	}

	return c.Status(status).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	_, err := database.DB.NewUpdate().Model(fact).Where("id = ?", id).Exec(c.Context())
	if err != nil {
		return c.Status(200).JSON(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	_, err := database.DB.NewDelete().Model(fact).Where("id = ?", id).Exec(c.Context())
	if err != nil {
		return NotFoundMessageJson(c, "")
	}

	return SuccessMessageJson(c, "Запись удалена")
}
