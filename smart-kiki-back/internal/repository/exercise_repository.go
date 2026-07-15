package repository

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type ExerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

func (r *ExerciseRepository) List(muscleGroup string) ([]model.Exercise, error) {
	var exercises []model.Exercise
	query := r.db.Order("muscle_group, name")
	if muscleGroup != "" {
		query = query.Where("muscle_group = ?", muscleGroup)
	}
	err := query.Find(&exercises).Error
	return exercises, err
}

func (r *ExerciseRepository) ExistsAll(ids []uuid.UUID) (bool, error) {
	if len(ids) == 0 {
		return true, nil
	}
	var count int64
	if err := r.db.Model(&model.Exercise{}).Where("id IN ?", ids).Count(&count).Error; err != nil {
		return false, err
	}
	return int(count) == len(ids), nil
}
