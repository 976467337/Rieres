package repository

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type AssessmentRepository struct {
	db *gorm.DB
}

func NewAssessmentRepository(db *gorm.DB) *AssessmentRepository {
	return &AssessmentRepository{db: db}
}

func (r *AssessmentRepository) Create(assessment *model.Assessment) error {
	return r.db.Create(assessment).Error
}

func (r *AssessmentRepository) ListByStudent(studentID uuid.UUID) ([]model.Assessment, error) {
	var assessments []model.Assessment
	err := r.db.Where("student_id = ?", studentID).Order("recorded_at DESC").Find(&assessments).Error
	return assessments, err
}
