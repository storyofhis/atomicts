package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/storyofhis/atomicts/atomicts-be/database"
	"github.com/storyofhis/atomicts/atomicts-be/database/models"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(`cannot load your env`, err.Error())
		return err
	}
	database.ConnectDB()
	return models.DB.AutoMigrate(&models.User{})
}