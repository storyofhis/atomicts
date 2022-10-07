package database

import (
	"fmt"
	"log"
	"os"

	"github.com/storyofhis/atomicts/atomicts-be/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_sslmode := os.Getenv("DB_SLLMODE")

	var err error
	dsn := fmt.Sprintf(
		`user=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		db_host, db_port, db_user, db_pass, db_name, db_sslmode,
	)

	models.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(`cannot open your database`)
		return
	}
}

func CloseConnection(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		log.Println(`cannot close connection database`)
		return
	}
	DB.Close()
}
