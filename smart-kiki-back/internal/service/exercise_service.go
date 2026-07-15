package service

import (
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

type ExerciseService struct {
	exerciseRepo *repository.ExerciseRepository
}

func NewExerciseService(exerciseRepo *repository.ExerciseRepository) *ExerciseService {
	return &ExerciseService{exerciseRepo: exerciseRepo}
}

func (s *ExerciseService) List(muscleGroup string) ([]model.Exercise, error) {
	return s.exerciseRepo.List(muscleGroup)
}
