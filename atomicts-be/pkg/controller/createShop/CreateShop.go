package createShop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"test" : "ccreate",
	})
}