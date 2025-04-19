package repositories

import "fmt"

type MessageRepository interface {
	StoreMessage() string
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) StoreMessage() string {
	return fmt.Sprintf("Message stored successfully, this is message repository")
}
