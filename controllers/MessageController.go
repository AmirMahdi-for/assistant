package controllers

import (
	"assistant/helpers"
	"assistant/models"
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
	userId := 1

	body, ok := helpers.ParseJSONBody(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var con *models.Conversation
	var err error

	if body["conversationId"].(float64) == 0 {
		con, err = mc.ConvRepo.CreateConversation(body, userId)
	} else {
		con, err = mc.ConvRepo.GetConversationById(body, userId)
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	result := mc.MessageRepo.StoreMessage()

	c.JSON(http.StatusOK, gin.H{
		"message":      result,
		"conversation": con,
	})
}
