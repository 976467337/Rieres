package repository

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *model.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) ListConversation(trainerID, studentID uuid.UUID) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.
		Where("trainer_id = ? AND student_id = ?", trainerID, studentID).
		Order("created_at").
		Find(&messages).Error
	return messages, err
}
