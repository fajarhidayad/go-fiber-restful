package user

import "github.com/gofiber/fiber/v2"

func UserRoutes(router fiber.Router) {
	r := router.Group("/user")

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User found",
		})
	})
}
