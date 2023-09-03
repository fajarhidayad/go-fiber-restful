package main

import (
	"log"

	"github.com/fajarhidayad/go-fiber-restful/db"
	"github.com/fajarhidayad/go-fiber-restful/routes"
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

	routes.Routes(app)

	log.Fatal(app.Listen(":8000"))
}
