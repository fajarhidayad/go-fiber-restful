package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/fajarhidayad/go-fiber-restful/db"
	"github.com/fajarhidayad/go-fiber-restful/handlers"
	"github.com/fajarhidayad/go-fiber-restful/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSignIn struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"-"`
}

func GetAllUsers(c *fiber.Ctx) error {
	db := db.DB
	var users []models.User

	db.Find(&users)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User data found",
		"data":    users,
	})

}

var Validation = validator.New()

func SignUp(c *fiber.Ctx) error {
	db := db.DB
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return handlers.CustomError(c, http.StatusBadRequest, "Invalid request")
	}

	if _, err := govalidator.ValidateStruct(&user); err != nil {
		var errs []string
		for _, msg := range err.(govalidator.Errors).Errors() {
			errs = append(errs, msg.Error())
		}
		return handlers.CustomError(c, http.StatusBadRequest, strings.Join(errs, " and "))
	}

	var checkUser models.User
	if err := db.Where("email = ?", user.Email).First(&checkUser).Error; err == nil {
		return handlers.CustomError(c, http.StatusBadRequest, "User already exist")
	}

	user.Password = handlers.HashPassword(user.Password)
	db.Create(&user)

	token := handlers.SignToken(user.Name, user.Email)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success sign up",
		"token":   fmt.Sprintf("Bearer %s", token),
	})

}

func SignIn(c *fiber.Ctx) error {
	db := db.DB
	var user UserSignIn
	if err := c.BodyParser(&user); err != nil {
		return handlers.CustomError(c, http.StatusBadRequest, "Invalid request")
	}

	if _, err := govalidator.ValidateStruct(&user); err != nil {
		var errs []string
		for _, msg := range err.(govalidator.Errors).Errors() {
			errs = append(errs, msg.Error())
		}
		return handlers.CustomError(c, http.StatusBadRequest, strings.Join(errs, " and "))
	}

	var result models.User

	if err := db.Where("email = ?", user.Email).First(&result).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return handlers.ErrorNotFound(c, "User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return handlers.CustomError(c, http.StatusBadRequest, "Password mismatch")
	}

	token := handlers.SignToken(result.Name, result.Email)

	return c.JSON(fiber.Map{
		"status": "success",
		"token":  fmt.Sprintf("Bearer %s", token),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := db.DB
	id, _ := c.ParamsInt("id")

	var user models.User
	if err := db.First(&user, uint(id)).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "Not Found",
			"message": "User not found",
		})
	}

	db.Delete(&user, id)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User deleted",
	})
}
