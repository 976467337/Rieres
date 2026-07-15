package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

var ErrWorkoutNotFound = errors.New("workout not found")

type WorkoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	return &WorkoutRepository{db: db}
}

func (r *WorkoutRepository) Create(workout *model.Workout) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Exercises").Create(workout).Error; err != nil {
			return err
		}
		for i := range workout.Exercises {
			workout.Exercises[i].ID = uuid.New()
			workout.Exercises[i].WorkoutID = workout.ID
			workout.Exercises[i].Position = i
		}
		if len(workout.Exercises) > 0 {
			if err := tx.Create(&workout.Exercises).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *WorkoutRepository) FindByID(id uuid.UUID) (*model.Workout, error) {
	var workout model.Workout
	err := r.db.
		Preload("Exercises", func(db *gorm.DB) *gorm.DB { return db.Order("position") }).
		Preload("Exercises.Exercise").
		Where("id = ?", id).
		First(&workout).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWorkoutNotFound
		}
		return nil, err
	}
	return &workout, nil
}

func (r *WorkoutRepository) ListByTrainer(trainerID uuid.UUID, studentID *uuid.UUID) ([]model.Workout, error) {
	var workouts []model.Workout
	query := r.db.Where("trainer_id = ?", trainerID).Order("created_at DESC")
	if studentID != nil {
		query = query.Where("student_id = ?", *studentID)
	}
	err := query.Find(&workouts).Error
	return workouts, err
}

func (r *WorkoutRepository) ListByStudent(studentID uuid.UUID) ([]model.Workout, error) {
	var workouts []model.Workout
	err := r.db.Where("student_id = ?", studentID).Order("created_at DESC").Find(&workouts).Error
	return workouts, err
}

func (r *WorkoutRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Workout{}, "id = ?", id).Error
}
