package model

import (
	"time"

	"github.com/google/uuid"
)

type Exercise struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	MuscleGroup  string    `gorm:"not null;index" json:"muscle_group"`
	Equipment    string    `json:"equipment"`
	Instructions string    `json:"instructions"`
	CreatedAt    time.Time `json:"created_at"`
}

func (Exercise) TableName() string {
	return "exercises"
}
