package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrNutritionPlanNotFound = errors.New("nutrition plan not found")

type NutritionPlanRepository struct {
	db *gorm.DB
}

func NewNutritionPlanRepository(db *gorm.DB) *NutritionPlanRepository {
	return &NutritionPlanRepository{db: db}
}

func (r *NutritionPlanRepository) Upsert(plan *model.NutritionPlan) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "trainer_id"}, {Name: "student_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"content", "updated_at"}),
	}).Create(plan).Error
}

func (r *NutritionPlanRepository) FindByTrainerAndStudent(trainerID, studentID uuid.UUID) (*model.NutritionPlan, error) {
	var plan model.NutritionPlan
	err := r.db.Where("trainer_id = ? AND student_id = ?", trainerID, studentID).First(&plan).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNutritionPlanNotFound
		}
		return nil, err
	}
	return &plan, nil
}

func (r *NutritionPlanRepository) FindLatestByStudent(studentID uuid.UUID) (*model.NutritionPlan, error) {
	var plan model.NutritionPlan
	err := r.db.Where("student_id = ?", studentID).Order("updated_at DESC").First(&plan).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNutritionPlanNotFound
		}
		return nil, err
	}
	return &plan, nil
}
