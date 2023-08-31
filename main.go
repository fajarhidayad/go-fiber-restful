package main

import (
	"log"

	"github.com/fajarhidayad/go-fiber-restful/db"
	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	app := fiber.New()

	// Connect to Postgres DB
	db.ConnectDB()

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It works")
	})

	log.Fatal(app.Listen(":8000"))
}
