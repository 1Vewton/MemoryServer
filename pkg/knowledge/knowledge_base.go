package knowledge

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// KnowledgeBase defines a knowledge base stored
type KnowledgeBase struct {
	gorm.Model
	ID          string
	Description string
	Name        string `gorm:"primaryKey"`
}

// NewKnowledgeBase creates new knowledge base
func NewKnowledgeBase() *KnowledgeBase {
	return &KnowledgeBase{
		ID: uuid.NewString(),
	}
}
