package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

var (
	ErrStudentLimitReached  = errors.New("student limit reached for current plan")
	ErrStudentNotFound      = errors.New("student not found")
	ErrStudentAlreadyLinked = errors.New("student already linked to this trainer")
)

type TrainerStudentService struct {
	trainerStudentRepo *repository.TrainerStudentRepository
	subscriptionRepo   *repository.SubscriptionRepository
	userRepo           *repository.UserRepository
}

func NewTrainerStudentService(
	trainerStudentRepo *repository.TrainerStudentRepository,
	subscriptionRepo *repository.SubscriptionRepository,
	userRepo *repository.UserRepository,
) *TrainerStudentService {
	return &TrainerStudentService{
		trainerStudentRepo: trainerStudentRepo,
		subscriptionRepo:   subscriptionRepo,
		userRepo:           userRepo,
	}
}

func (s *TrainerStudentService) AddStudent(trainerID uuid.UUID, email string) (*model.User, error) {
	student, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrStudentNotFound
		}
		return nil, err
	}

	if student.Role != model.RoleClient {
		return nil, ErrStudentNotFound
	}

	exists, err := s.trainerStudentRepo.ExistsByTrainerAndStudent(trainerID, student.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrStudentAlreadyLinked
	}

	sub, err := s.subscriptionRepo.FindByTrainerID(trainerID)
	if err != nil {
		return nil, err
	}

	count, err := s.trainerStudentRepo.CountByTrainer(trainerID)
	if err != nil {
		return nil, err
	}

	if limit := model.StudentLimit(sub.Plan); limit != nil && count >= int64(*limit) {
		return nil, ErrStudentLimitReached
	}

	link := &model.TrainerStudent{
		ID:        uuid.New(),
		TrainerID: trainerID,
		StudentID: student.ID,
	}
	if err := s.trainerStudentRepo.Create(link); err != nil {
		return nil, err
	}

	return student, nil
}

func (s *TrainerStudentService) List(trainerID uuid.UUID) ([]model.User, error) {
	return s.trainerStudentRepo.ListStudentsByTrainer(trainerID)
}
