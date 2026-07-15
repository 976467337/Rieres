package repository

import (
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

type TrainerStudentRepository struct {
	db *gorm.DB
}

func NewTrainerStudentRepository(db *gorm.DB) *TrainerStudentRepository {
	return &TrainerStudentRepository{db: db}
}

func (r *TrainerStudentRepository) Create(link *model.TrainerStudent) error {
	return r.db.Create(link).Error
}

func (r *TrainerStudentRepository) CountByTrainer(trainerID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&model.TrainerStudent{}).Where("trainer_id = ?", trainerID).Count(&count).Error
	return count, err
}

func (r *TrainerStudentRepository) ExistsByTrainerAndStudent(trainerID, studentID uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&model.TrainerStudent{}).
		Where("trainer_id = ? AND student_id = ?", trainerID, studentID).
		Count(&count).Error
	return count > 0, err
}

func (r *TrainerStudentRepository) ListStudentsByTrainer(trainerID uuid.UUID) ([]model.User, error) {
	var users []model.User
	err := r.db.
		Joins("JOIN trainer_students ON trainer_students.student_id = users.id").
		Where("trainer_students.trainer_id = ?", trainerID).
		Order("trainer_students.created_at").
		Find(&users).Error
	return users, err
}

func (r *TrainerStudentRepository) ListTrainersByStudent(studentID uuid.UUID) ([]model.User, error) {
	var users []model.User
	err := r.db.
		Joins("JOIN trainer_students ON trainer_students.trainer_id = users.id").
		Where("trainer_students.student_id = ?", studentID).
		Order("trainer_students.created_at").
		Find(&users).Error
	return users, err
}
