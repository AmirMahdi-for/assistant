package repositories

import "fmt"

type ConversationRepository interface {
	CreateConversation() string
}

type conversationRepository struct{}

func NewConversationRepository() ConversationRepository {
	return &conversationRepository{}
}

func (r *conversationRepository) CreateConversation() string {
	return fmt.Sprintf("this is conversation repository")
}
