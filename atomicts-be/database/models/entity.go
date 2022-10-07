package models

import "gorm.io/gorm"

var DB *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Username string
	Password string
}

var Body struct {
	Email    string
	Username string
	Password string
}
