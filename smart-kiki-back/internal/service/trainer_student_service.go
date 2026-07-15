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

	if err := s.link(trainerID, student.ID); err != nil {
		return nil, err
	}

	return student, nil
}

// RequestTrainer é o mesmo vínculo de AddStudent, mas iniciado pelo aluno a
// partir do marketplace (busca de personal), em vez do personal por e-mail.
func (s *TrainerStudentService) RequestTrainer(studentID, trainerID uuid.UUID) error {
	return s.link(trainerID, studentID)
}

func (s *TrainerStudentService) link(trainerID, studentID uuid.UUID) error {
	exists, err := s.trainerStudentRepo.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return err
	}
	if exists {
		return ErrStudentAlreadyLinked
	}

	sub, err := s.subscriptionRepo.FindByTrainerID(trainerID)
	if err != nil {
		return err
	}

	count, err := s.trainerStudentRepo.CountByTrainer(trainerID)
	if err != nil {
		return err
	}

	if limit := model.StudentLimit(sub.Plan); limit != nil && count >= int64(*limit) {
		return ErrStudentLimitReached
	}

	return s.trainerStudentRepo.Create(&model.TrainerStudent{
		ID:        uuid.New(),
		TrainerID: trainerID,
		StudentID: studentID,
	})
}

func (s *TrainerStudentService) List(trainerID uuid.UUID) ([]model.User, error) {
	return s.trainerStudentRepo.ListStudentsByTrainer(trainerID)
}

func (s *TrainerStudentService) ListTrainersForStudent(studentID uuid.UUID) ([]model.User, error) {
	return s.trainerStudentRepo.ListTrainersByStudent(studentID)
}
