package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func SuccessMessageJson(c *fiber.Ctx, msg string) error {
	message := Message{
		Message: msg,
		Success: true,
	}

	return c.Status(200).JSON(message)
}

func NotFoundMessageJson(c *fiber.Ctx, msg string) error {
	if msg == "" {
		msg = "Запись не найдена"
	}

	message := Message{
		Message: msg,
		Success: false,
	}

	return c.Status(fiber.StatusNotFound).JSON(message)
}

func SystemErrorMessageJson(c *fiber.Ctx, msg string) error {
	if msg == "" {
		msg = "Системная ошибка"
	}

	message := Message{
		Message: msg,
		Success: false,
	}

	return c.Status(fiber.StatusInternalServerError).JSON(message)
}

type ResponseRows struct {
	Count int         `json:"count"`
	Rows  interface{} `json:"rows"`
}
