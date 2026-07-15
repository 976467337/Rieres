package model

import (
	"time"

	"github.com/google/uuid"
)

type Assessment struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID  uuid.UUID `gorm:"type:uuid;not null;index" json:"trainer_id"`
	StudentID  uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	RecordedAt time.Time `gorm:"not null" json:"recorded_at"`
	WeightKg   *float64  `json:"weight_kg"`
	Notes      string    `json:"notes"`
	CreatedAt  time.Time `json:"created_at"`
}

func (Assessment) TableName() string {
	return "assessments"
}

type CreateAssessmentRequest struct {
	RecordedAt time.Time `json:"recorded_at" binding:"required"`
	WeightKg   *float64  `json:"weight_kg"`
	Notes      string    `json:"notes"`
}
