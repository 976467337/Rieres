package model

import (
	"time"

	"github.com/google/uuid"
)

type TrainerStudent struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID uuid.UUID `gorm:"type:uuid;not null;index" json:"trainer_id"`
	StudentID uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (TrainerStudent) TableName() string {
	return "trainer_students"
}

type AddStudentRequest struct {
	Email string `json:"email" binding:"required,email"`
}
