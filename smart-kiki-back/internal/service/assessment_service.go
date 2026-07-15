package service

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

type AssessmentService struct {
	assessmentRepo *repository.AssessmentRepository
	trainerStudent *repository.TrainerStudentRepository
}

func NewAssessmentService(assessmentRepo *repository.AssessmentRepository, trainerStudentRepo *repository.TrainerStudentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo, trainerStudent: trainerStudentRepo}
}

func (s *AssessmentService) Create(trainerID, studentID uuid.UUID, req *model.CreateAssessmentRequest) (*model.Assessment, error) {
	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	assessment := &model.Assessment{
		ID:         uuid.New(),
		TrainerID:  trainerID,
		StudentID:  studentID,
		RecordedAt: req.RecordedAt,
		WeightKg:   req.WeightKg,
		Notes:      req.Notes,
	}
	if err := s.assessmentRepo.Create(assessment); err != nil {
		return nil, err
	}
	return assessment, nil
}

func (s *AssessmentService) ListForStudent(studentID uuid.UUID) ([]model.Assessment, error) {
	return s.assessmentRepo.ListByStudent(studentID)
}
