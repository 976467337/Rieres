package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

var ErrSessionForbidden = errors.New("session does not belong to this user")

type SessionService struct {
	sessionRepo    *repository.SessionRepository
	trainerStudent *repository.TrainerStudentRepository
}

func NewSessionService(sessionRepo *repository.SessionRepository, trainerStudentRepo *repository.TrainerStudentRepository) *SessionService {
	return &SessionService{sessionRepo: sessionRepo, trainerStudent: trainerStudentRepo}
}

func (s *SessionService) Create(trainerID uuid.UUID, req *model.CreateSessionRequest) (*model.Session, error) {
	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, req.StudentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	duration := req.DurationMinutes
	if duration <= 0 {
		duration = 60
	}

	session := &model.Session{
		ID:              uuid.New(),
		TrainerID:       trainerID,
		StudentID:       req.StudentID,
		ScheduledAt:     req.ScheduledAt,
		DurationMinutes: duration,
		Notes:           req.Notes,
		Status:          model.SessionStatusScheduled,
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) ListForTrainer(trainerID uuid.UUID) ([]model.Session, error) {
	return s.sessionRepo.ListByTrainer(trainerID)
}

func (s *SessionService) ListForStudent(studentID uuid.UUID) ([]model.Session, error) {
	return s.sessionRepo.ListByStudent(studentID)
}

func (s *SessionService) Cancel(userID, sessionID uuid.UUID) error {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return err
	}
	if session.TrainerID != userID && session.StudentID != userID {
		return ErrSessionForbidden
	}
	return s.sessionRepo.UpdateStatus(sessionID, model.SessionStatusCancelled)
}
