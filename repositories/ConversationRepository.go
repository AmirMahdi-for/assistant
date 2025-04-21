package repositories

import (
	"assistant/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ConversationRepository interface {
	CreateConversation(body map[string]interface{}, userId int) (*models.Conversation, error)
	GetConversationById(body map[string]interface{}, userId int) (*models.Conversation, error)
}

type conversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) ConversationRepository {
	return &conversationRepository{db: db}
}

func (r *conversationRepository) CreateConversation(body map[string]interface{}, userId int) (*models.Conversation, error) {
	var conv models.Conversation
	fmt.Println("Creating conversation with body:", body)
	messageRaw, ok := body["message"].(map[string]interface{})
	if !ok {
		return nil, errors.New("'message' is missing or not in the expected format")
	}

	textRaw, ok := messageRaw["text"].(string)
	if !ok || textRaw == "" {
		return nil, errors.New("text is required and must be a non-empty string")
	}

	runes := []rune(textRaw)
	if len(runes) <= 20 {
		conv.Title = string(runes)
	} else {
		conv.Title = string(runes[:20])
	}

	conv.UserID = userId

	if err := r.db.Create(&conv).Error; err != nil {
		return nil, err
	}

	return &conv, nil
}

func (r *conversationRepository) GetConversationById(body map[string]interface{}, userId int) (*models.Conversation, error) {
	var conv models.Conversation
	conversationId, ok := body["conversationId"].(float64)

	if !ok {
		return nil, errors.New("'conversationId' is missing or not in the expected format")
	}
	fmt.Println("Getting conversation with id:", conversationId)
	if err := r.db.First(&conv, conversationId).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("conversation not found")
		}
		return nil, err
	}

	return &conv, nil
}
