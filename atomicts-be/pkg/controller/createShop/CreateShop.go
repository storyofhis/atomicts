package createShop

import (
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/storyofhis/atomicts/atomicts-be/database/models"
)

type Avatar struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type User struct {
	User   models.User
	Avatar Avatar
}

func CreateController(c *gin.Context) {
	var userObj User
	if err := c.ShouldBind(&userObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}
	if err := c.ShouldBindUri(&userObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "bad request",
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is Received",
		})
		return
	}
	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "/some/path/on/server"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "uploaded successfully",
	})
}
