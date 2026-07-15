package service

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

type MessageService struct {
	messageRepo    *repository.MessageRepository
	trainerStudent *repository.TrainerStudentRepository
}

func NewMessageService(messageRepo *repository.MessageRepository, trainerStudentRepo *repository.TrainerStudentRepository) *MessageService {
	return &MessageService{messageRepo: messageRepo, trainerStudent: trainerStudentRepo}
}

func resolveConversation(userID uuid.UUID, role string, otherUserID uuid.UUID) (trainerID, studentID uuid.UUID) {
	if role == string(model.RoleTrainer) {
		return userID, otherUserID
	}
	return otherUserID, userID
}

func (s *MessageService) Send(senderID uuid.UUID, senderRole string, otherUserID uuid.UUID, body string) (*model.Message, error) {
	trainerID, studentID := resolveConversation(senderID, senderRole, otherUserID)

	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	message := &model.Message{
		ID:        uuid.New(),
		TrainerID: trainerID,
		StudentID: studentID,
		SenderID:  senderID,
		Body:      body,
	}
	if err := s.messageRepo.Create(message); err != nil {
		return nil, err
	}
	return message, nil
}

func (s *MessageService) ListConversation(userID uuid.UUID, role string, otherUserID uuid.UUID) ([]model.Message, error) {
	trainerID, studentID := resolveConversation(userID, role, otherUserID)

	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	return s.messageRepo.ListConversation(trainerID, studentID)
}
