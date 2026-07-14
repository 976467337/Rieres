package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

var ErrInvalidPlan = errors.New("invalid plan")

type SubscriptionService struct {
	subscriptionRepo *repository.SubscriptionRepository
	trainerStudent   *repository.TrainerStudentRepository
}

func NewSubscriptionService(subscriptionRepo *repository.SubscriptionRepository, trainerStudentRepo *repository.TrainerStudentRepository) *SubscriptionService {
	return &SubscriptionService{subscriptionRepo: subscriptionRepo, trainerStudent: trainerStudentRepo}
}

func (s *SubscriptionService) EnsureFreeSubscription(trainerID uuid.UUID) error {
	if _, err := s.subscriptionRepo.FindByTrainerID(trainerID); err == nil {
		return nil
	} else if !errors.Is(err, repository.ErrSubscriptionNotFound) {
		return err
	}

	return s.subscriptionRepo.Create(&model.TrainerSubscription{
		TrainerID:          trainerID,
		Plan:               model.PlanFree,
		CurrentPeriodStart: time.Now(),
	})
}

func (s *SubscriptionService) Get(trainerID uuid.UUID) (*model.SubscriptionResponse, error) {
	sub, err := s.subscriptionRepo.FindByTrainerID(trainerID)
	if err != nil {
		return nil, err
	}

	count, err := s.trainerStudent.CountByTrainer(trainerID)
	if err != nil {
		return nil, err
	}

	return s.toResponse(sub, count), nil
}

func (s *SubscriptionService) ChangePlan(trainerID uuid.UUID, plan model.PlanType) (*model.SubscriptionResponse, error) {
	if !model.IsValidPlan(plan) {
		return nil, ErrInvalidPlan
	}

	sub, err := s.subscriptionRepo.FindByTrainerID(trainerID)
	if err != nil {
		return nil, err
	}

	sub.Plan = plan
	sub.CurrentPeriodStart = time.Now()
	if err := s.subscriptionRepo.Update(sub); err != nil {
		return nil, err
	}

	count, err := s.trainerStudent.CountByTrainer(trainerID)
	if err != nil {
		return nil, err
	}

	return s.toResponse(sub, count), nil
}

func (s *SubscriptionService) toResponse(sub *model.TrainerSubscription, studentsCount int64) *model.SubscriptionResponse {
	return &model.SubscriptionResponse{
		Plan:                 string(sub.Plan),
		Price:                model.PlanPrice(sub.Plan),
		StudentLimit:         model.StudentLimit(sub.Plan),
		StudentsCount:        studentsCount,
		VisibleInMarketplace: model.IsVisibleInMarketplace(*sub),
		CurrentPeriodStart:   sub.CurrentPeriodStart,
	}
}
