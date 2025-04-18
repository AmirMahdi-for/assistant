package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Store(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "hello agha amirmahdi",
	})
}
