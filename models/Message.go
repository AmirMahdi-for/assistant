package models

import (
	"gorm.io/gorm"
)

type Message struct {
	ID             uint            `gorm:"primaryKey"`
	UserID         int             `gorm:"not null"`
	ConversationID int             `gorm:"not null"`
	Message        string          `gorm:"type:json;default:null"`
	ReplayID       int             `gorm:"default:0"`
	From           string          `gorm:"type:enum('user', 'server');default:'user'"`
	Action         string          `gorm:"type:varchar(255);default:null"`
	CreatedAt      *gorm.DeletedAt `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      *gorm.DeletedAt `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt      *gorm.DeletedAt `gorm:"index"` // For Soft Delete
}
