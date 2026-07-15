package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(session *model.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) FindByID(id uuid.UUID) (*model.Session, error) {
	var session model.Session
	if err := r.db.Where("id = ?", id).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) ListByTrainer(trainerID uuid.UUID) ([]model.Session, error) {
	var sessions []model.Session
	err := r.db.Where("trainer_id = ?", trainerID).Order("scheduled_at").Find(&sessions).Error
	return sessions, err
}

func (r *SessionRepository) ListByStudent(studentID uuid.UUID) ([]model.Session, error) {
	var sessions []model.Session
	err := r.db.Where("student_id = ?", studentID).Order("scheduled_at").Find(&sessions).Error
	return sessions, err
}

func (r *SessionRepository) UpdateStatus(id uuid.UUID, status string) error {
	return r.db.Model(&model.Session{}).Where("id = ?", id).Update("status", status).Error
}
