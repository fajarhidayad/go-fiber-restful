package handlers

import (
	"time"

	"github.com/fajarhidayad/go-fiber-restful/config"
	"github.com/golang-jwt/jwt/v5"
)

func SignToken(name string, email string) string {
	claims := jwt.MapClaims{
		"name":  name,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.Config("SECRET_KEY")
	t, _ := token.SignedString([]byte(secret))

	return t
}
