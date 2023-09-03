package routes

import (
	"github.com/fajarhidayad/go-fiber-restful/routes/product"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1", logger.New())

	product.ProductRoutes(v1)
}
