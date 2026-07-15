package model

import (
	"time"

	"github.com/google/uuid"
)

type Workout struct {
	ID        uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	TrainerID uuid.UUID         `gorm:"type:uuid;not null;index" json:"trainer_id"`
	StudentID uuid.UUID         `gorm:"type:uuid;not null;index" json:"student_id"`
	Name      string            `gorm:"not null" json:"name"`
	Notes     string            `json:"notes"`
	CreatedAt time.Time         `json:"created_at"`
	Exercises []WorkoutExercise `gorm:"foreignKey:WorkoutID" json:"exercises,omitempty"`
}

func (Workout) TableName() string {
	return "workouts"
}

type WorkoutExercise struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WorkoutID   uuid.UUID `gorm:"type:uuid;not null;index" json:"workout_id"`
	ExerciseID  uuid.UUID `gorm:"type:uuid;not null" json:"exercise_id"`
	Position    int       `gorm:"not null" json:"position"`
	Sets        int       `gorm:"not null" json:"sets"`
	Reps        int       `gorm:"not null" json:"reps"`
	Load        string    `json:"load"`
	RestSeconds int       `json:"rest_seconds"`
	Notes       string    `json:"notes"`
	Exercise    Exercise  `gorm:"foreignKey:ExerciseID" json:"exercise"`
}

func (WorkoutExercise) TableName() string {
	return "workout_exercises"
}

type WorkoutLog struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	WorkoutID   uuid.UUID `gorm:"type:uuid;not null;index" json:"workout_id"`
	StudentID   uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	CompletedAt time.Time `json:"completed_at"`
}

func (WorkoutLog) TableName() string {
	return "workout_logs"
}

type WorkoutExerciseRequest struct {
	ExerciseID  uuid.UUID `json:"exercise_id" binding:"required"`
	Sets        int       `json:"sets" binding:"required,min=1"`
	Reps        int       `json:"reps" binding:"required,min=1"`
	Load        string    `json:"load"`
	RestSeconds int       `json:"rest_seconds"`
	Notes       string    `json:"notes"`
}

type CreateWorkoutRequest struct {
	StudentID uuid.UUID                `json:"student_id" binding:"required"`
	Name      string                   `json:"name" binding:"required"`
	Notes     string                   `json:"notes"`
	Exercises []WorkoutExerciseRequest `json:"exercises" binding:"required,min=1,dive"`
}
