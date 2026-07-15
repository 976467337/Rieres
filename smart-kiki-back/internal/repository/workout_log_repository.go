package repository

import (
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type WorkoutLogRepository struct {
	db *gorm.DB
}

func NewWorkoutLogRepository(db *gorm.DB) *WorkoutLogRepository {
	return &WorkoutLogRepository{db: db}
}

func (r *WorkoutLogRepository) Create(log *model.WorkoutLog) error {
	return r.db.Create(log).Error
}
