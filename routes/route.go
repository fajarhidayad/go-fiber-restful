package routes

import (
	"github.com/fajarhidayad/go-fiber-restful/handlers"
	"github.com/fajarhidayad/go-fiber-restful/routes/product"
	"github.com/fajarhidayad/go-fiber-restful/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1", logger.New())

	product.ProductRoutes(v1)
	user.UserRoutes(v1)

	app.Use(func(c *fiber.Ctx) error {
		return handlers.ErrorNotFound(c, "Route Not Found")
	})
}
