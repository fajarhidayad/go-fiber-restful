package controllers

import (
	"net/http"

	"github.com/fajarhidayad/go-fiber-restful/db"
	"github.com/fajarhidayad/go-fiber-restful/handlers"
	"github.com/fajarhidayad/go-fiber-restful/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	db := db.DB
	var products []models.Product

	db.Find(&products)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	db := db.DB
	id, _ := c.ParamsInt("id")
	var product models.Product

	if err := db.First(&product, uint(id)).Error; err != nil {
		return handlers.ErrorNotFound(c, "Product not found")
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    product,
	})
}

func CreateNewProduct(c *fiber.Ctx) error {
	db := db.DB

	newProduct := new(models.Product)

	if err := c.BodyParser(&newProduct); err != nil {
		return handlers.CustomError(c, http.StatusBadRequest, "Please fill the form")
	}
	db.Create(&newProduct)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Created new product",
	})
}

func ChangeProductById(c *fiber.Ctx) error {
	db := db.DB

	updatedProduct := new(models.Product)

	if err := c.BodyParser(&updatedProduct); err != nil {
		return handlers.CustomError(c, http.StatusBadRequest, "Please fill the form")
	}

	id, _ := c.ParamsInt("id")

	var result models.Product
	if err := db.First(&result, uint(id)).Error; err != nil {
		return handlers.ErrorNotFound(c, "Product not found")
	}

	db.Model(&result).Updates(updatedProduct)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	db := db.DB

	var product models.Product
	id, _ := c.ParamsInt("id")

	if err := db.First(&product, uint(id)).Error; err != nil {
		return handlers.ErrorNotFound(c, "Product not found")
	}

	db.Delete(&product, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product deleted",
	})
}
