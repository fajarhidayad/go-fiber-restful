package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
