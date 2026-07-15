package service

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

type NutritionPlanService struct {
	nutritionPlanRepo *repository.NutritionPlanRepository
	trainerStudent    *repository.TrainerStudentRepository
}

func NewNutritionPlanService(nutritionPlanRepo *repository.NutritionPlanRepository, trainerStudentRepo *repository.TrainerStudentRepository) *NutritionPlanService {
	return &NutritionPlanService{nutritionPlanRepo: nutritionPlanRepo, trainerStudent: trainerStudentRepo}
}

func (s *NutritionPlanService) Upsert(trainerID, studentID uuid.UUID, content string) (*model.NutritionPlan, error) {
	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	plan := &model.NutritionPlan{TrainerID: trainerID, StudentID: studentID, Content: content}
	if err := s.nutritionPlanRepo.Upsert(plan); err != nil {
		return nil, err
	}
	return s.nutritionPlanRepo.FindByTrainerAndStudent(trainerID, studentID)
}

func (s *NutritionPlanService) GetForTrainerAndStudent(trainerID, studentID uuid.UUID) (*model.NutritionPlan, error) {
	return s.nutritionPlanRepo.FindByTrainerAndStudent(trainerID, studentID)
}

func (s *NutritionPlanService) GetLatestForStudent(studentID uuid.UUID) (*model.NutritionPlan, error) {
	return s.nutritionPlanRepo.FindLatestByStudent(studentID)
}
