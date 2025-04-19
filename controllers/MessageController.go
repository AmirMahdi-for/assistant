package controllers

import (
	"assistant/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {
	repo repositories.MessageRepository
}

func NewMessageController(repo repositories.MessageRepository) *MessageController {
	return &MessageController{repo: repo}
}

func (mc *MessageController) Store(c *gin.Context) {
	result := mc.repo.StoreMessage()
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
