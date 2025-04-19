package models

import (
	"gorm.io/gorm"
)

type Conversation struct {
	ID        uint            `gorm:"primaryKey"`
	UserID    int             `gorm:"not null"`
	Title     string          `gorm:"not null"`
	Stage     string          `gorm:"type:enum('active', 'deactive');default:'active'"`
	Counter   int             `gorm:"default:1"`
	CreatedAt *gorm.DeletedAt `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt *gorm.DeletedAt `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `gorm:"index"` // For Soft Delete
}
