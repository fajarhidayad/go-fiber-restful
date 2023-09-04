package user

import (
	"github.com/fajarhidayad/go-fiber-restful/controllers"
	"github.com/fajarhidayad/go-fiber-restful/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	r := router.Group("/user")

	r.Get("/", controllers.GetAllUsers)
	r.Post("/signup", controllers.SignUp)
	r.Post("/signin", controllers.SignIn)

	r.Use(middlewares.JWTMiddleware)
	r.Delete("/:id", controllers.DeleteUser)
}
