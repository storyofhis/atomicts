package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/storyofhis/atomicts/atomicts-be/database/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUpController(c *gin.Context) {
	if c.Bind(&models.Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(models.Body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := &models.User{
		Email:    models.Body.Email,
		Username: models.Body.Username,
		Password: string(hash),
	}

	result := models.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success to create user",
		// "payload": {
		// 	"username" : user.Username,
		// 	"email" : user.Email,
		// },
	})
}

func LoginController(c *gin.Context) {
	if c.Bind(&models.Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errur": "Failed to read body",
		})
	}

	var user models.User
	if err := models.DB.First(&user, "email = ?", models.Body.Email).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found",
		})
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(models.Body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password doesn't mathch",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"status": "success to login",
	})
}

func ValidateController(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"status": "I'm Logging in",
		"user":   user,
	})
}

func GetUserByID (c *gin.Context) {
	var User models.User
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "id cannot be empty",
		})
	}
	err := models.DB.Where("id = ?", id).First(&User).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : "user not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : "success found",
		"user" : User,
	})
}

func GetAllUser (c *gin.Context) {
	Users := &[]models.User{}
	if err := models.DB.Find(Users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"users" : Users,
	})
}	