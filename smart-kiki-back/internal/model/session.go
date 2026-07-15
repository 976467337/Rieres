package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	SessionStatusScheduled = "agendada"
	SessionStatusCancelled = "cancelada"
)

type Session struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID       uuid.UUID `gorm:"type:uuid;not null;index" json:"trainer_id"`
	StudentID       uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	ScheduledAt     time.Time `gorm:"not null" json:"scheduled_at"`
	DurationMinutes int       `gorm:"not null;default:60" json:"duration_minutes"`
	Notes           string    `json:"notes"`
	Status          string    `gorm:"type:varchar(20);not null;default:agendada" json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

func (Session) TableName() string {
	return "sessions"
}

type CreateSessionRequest struct {
	StudentID       uuid.UUID `json:"student_id" binding:"required"`
	ScheduledAt     time.Time `json:"scheduled_at" binding:"required"`
	DurationMinutes int       `json:"duration_minutes"`
	Notes           string    `json:"notes"`
}
