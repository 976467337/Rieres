package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID uuid.UUID `gorm:"type:uuid;not null;index" json:"trainer_id"`
	StudentID uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	SenderID  uuid.UUID `gorm:"type:uuid;not null" json:"sender_id"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (Message) TableName() string {
	return "messages"
}

type SendMessageRequest struct {
	ToUserID uuid.UUID `json:"to_user_id" binding:"required"`
	Body     string    `json:"body" binding:"required"`
}
