package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

var (
	ErrStudentNotLinked = errors.New("student is not linked to this trainer")
	ErrExerciseNotFound = errors.New("one or more exercises not found")
	ErrWorkoutForbidden = errors.New("workout does not belong to this user")
)

type WorkoutService struct {
	workoutRepo    *repository.WorkoutRepository
	workoutLogRepo *repository.WorkoutLogRepository
	exerciseRepo   *repository.ExerciseRepository
	trainerStudent *repository.TrainerStudentRepository
}

func NewWorkoutService(
	workoutRepo *repository.WorkoutRepository,
	workoutLogRepo *repository.WorkoutLogRepository,
	exerciseRepo *repository.ExerciseRepository,
	trainerStudentRepo *repository.TrainerStudentRepository,
) *WorkoutService {
	return &WorkoutService{
		workoutRepo:    workoutRepo,
		workoutLogRepo: workoutLogRepo,
		exerciseRepo:   exerciseRepo,
		trainerStudent: trainerStudentRepo,
	}
}

func (s *WorkoutService) Create(trainerID uuid.UUID, req *model.CreateWorkoutRequest) (*model.Workout, error) {
	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, req.StudentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrStudentNotLinked
	}

	exerciseIDs := make([]uuid.UUID, len(req.Exercises))
	items := make([]model.WorkoutExercise, len(req.Exercises))
	for i, ex := range req.Exercises {
		exerciseIDs[i] = ex.ExerciseID
		items[i] = model.WorkoutExercise{
			ExerciseID:  ex.ExerciseID,
			Sets:        ex.Sets,
			Reps:        ex.Reps,
			Load:        ex.Load,
			RestSeconds: ex.RestSeconds,
			Notes:       ex.Notes,
		}
	}

	exist, err := s.exerciseRepo.ExistsAll(exerciseIDs)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrExerciseNotFound
	}

	workout := &model.Workout{
		ID:        uuid.New(),
		TrainerID: trainerID,
		StudentID: req.StudentID,
		Name:      req.Name,
		Notes:     req.Notes,
		Exercises: items,
	}

	if err := s.workoutRepo.Create(workout); err != nil {
		return nil, err
	}

	return s.workoutRepo.FindByID(workout.ID)
}

func (s *WorkoutService) Get(id uuid.UUID) (*model.Workout, error) {
	return s.workoutRepo.FindByID(id)
}

func (s *WorkoutService) ListForTrainer(trainerID uuid.UUID, studentID *uuid.UUID) ([]model.Workout, error) {
	return s.workoutRepo.ListByTrainer(trainerID, studentID)
}

func (s *WorkoutService) ListForStudent(studentID uuid.UUID) ([]model.Workout, error) {
	return s.workoutRepo.ListByStudent(studentID)
}

func (s *WorkoutService) Delete(trainerID, workoutID uuid.UUID) error {
	workout, err := s.workoutRepo.FindByID(workoutID)
	if err != nil {
		return err
	}
	if workout.TrainerID != trainerID {
		return ErrWorkoutForbidden
	}
	return s.workoutRepo.Delete(workoutID)
}

func (s *WorkoutService) Complete(studentID, workoutID uuid.UUID) error {
	workout, err := s.workoutRepo.FindByID(workoutID)
	if err != nil {
		return err
	}
	if workout.StudentID != studentID {
		return ErrWorkoutForbidden
	}
	return s.workoutLogRepo.Create(&model.WorkoutLog{
		ID:          uuid.New(),
		WorkoutID:   workoutID,
		StudentID:   studentID,
		CompletedAt: time.Now(),
	})
}
