package model

import (
	"time"

	"github.com/google/uuid"
)

type NutritionPlan struct {
	TrainerID uuid.UUID `gorm:"type:uuid;primaryKey" json:"trainer_id"`
	StudentID uuid.UUID `gorm:"type:uuid;primaryKey" json:"student_id"`
	Content   string    `gorm:"not null" json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (NutritionPlan) TableName() string {
	return "nutrition_plans"
}

type UpsertNutritionPlanRequest struct {
	Content string `json:"content" binding:"required"`
}
