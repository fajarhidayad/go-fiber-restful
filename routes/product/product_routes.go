package product

import (
	"github.com/fajarhidayad/go-fiber-restful/controllers"
	"github.com/fajarhidayad/go-fiber-restful/middlewares"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(router fiber.Router) {
	r := router.Group("/product")

	r.Get("/", controllers.GetAllProduct)
	r.Get("/:id", controllers.GetProductById)

	r.Use(middlewares.JWTMiddleware)
	r.Post("/", controllers.CreateNewProduct)
	r.Put("/:id", controllers.ChangeProductById)
	r.Delete("/:id", controllers.DeleteProduct)
}
