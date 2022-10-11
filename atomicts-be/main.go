package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/storyofhis/atomicts/atomicts-be/database"
	"github.com/storyofhis/atomicts/atomicts-be/database/models"
	"github.com/storyofhis/atomicts/atomicts-be/pkg/controller"
	"github.com/storyofhis/atomicts/atomicts-be/pkg/controller/createShop"
	"github.com/storyofhis/atomicts/atomicts-be/pkg/middleware"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(`cannot loadf your env`, err)
		return
	}
	database.ConnectDB()
	models.DB.AutoMigrate(&models.User{})
}
func main() {
	router := gin.Default()

	// Authenticate
	v1 := router.Group("/v1")
	v1.GET("/reyhan", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "begitu syulit lupakan reyhan",
		})
	})
	v1.POST("/signup", controller.SignUpController)
	v1.POST("/login", controller.LoginController)
	v1.GET("/validate", middleware.RequireAuth, controller.ValidateController)

	// CREATE
	v1.GET("/validator/create-shop", middleware.RequireAuth, createShop.CreateController)

	router.Run(":8080")
}
