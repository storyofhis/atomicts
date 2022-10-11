package models

import (
	"mime/multipart"

	"gorm.io/gorm"
)

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

// products
type Products struct {
	gorm.Model
	Name  string
	Desc  string
	Price int
	User  User
}

type UploadFile struct {
	Product Products
	File    multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}
