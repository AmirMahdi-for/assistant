package controllers

import (
	"assistant/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {
	MessageRepo repositories.MessageRepository
	ConvRepo    repositories.ConversationRepository
}

func NewMessageController(msgRepo repositories.MessageRepository, repository repositories.ConversationRepository) *MessageController {
	return &MessageController{MessageRepo: msgRepo, ConvRepo: repository}
}

func (mc *MessageController) Store(c *gin.Context) {
	result := mc.MessageRepo.StoreMessage()
	con := mc.ConvRepo.CreateConversation()
	c.JSON(http.StatusOK, gin.H{
		"message":      result,
		"conversation": con,
	})
}
