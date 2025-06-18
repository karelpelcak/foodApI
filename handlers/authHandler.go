package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Yep its me",
	})
}
