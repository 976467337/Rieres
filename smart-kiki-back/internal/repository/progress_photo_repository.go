package repository

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type ProgressPhotoRepository struct {
	db *gorm.DB
}

func NewProgressPhotoRepository(db *gorm.DB) *ProgressPhotoRepository {
	return &ProgressPhotoRepository{db: db}
}

func (r *ProgressPhotoRepository) Create(photo *model.ProgressPhoto) error {
	return r.db.Create(photo).Error
}

func (r *ProgressPhotoRepository) ListByStudent(studentID uuid.UUID) ([]model.ProgressPhoto, error) {
	var photos []model.ProgressPhoto
	err := r.db.Where("student_id = ?", studentID).Order("recorded_at").Find(&photos).Error
	return photos, err
}
