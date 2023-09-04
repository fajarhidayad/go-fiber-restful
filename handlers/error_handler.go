package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorNotFound(c *fiber.Ctx, message string) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"status":  "failed",
		"message": message,
	})
}

func RouteNotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"status":  "not found",
		"message": "Route not found",
	})
}

func CustomError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
