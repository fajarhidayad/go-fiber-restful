package handlers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	cost := bcrypt.DefaultCost
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(encrypted)
}
