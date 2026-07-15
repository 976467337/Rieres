package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
	"github.com/smartkiki/api/internal/service"
	"github.com/smartkiki/api/pkg/middleware"
)

type WorkoutHandler struct {
	workoutService *service.WorkoutService
}

func NewWorkoutHandler(workoutService *service.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{workoutService: workoutService}
}

func (h *WorkoutHandler) Create(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	if role, _ := c.Get(middleware.ContextRoleKey); role != string(model.RoleTrainer) {
		c.JSON(http.StatusForbidden, gin.H{"error": "only trainers can create workouts"})
		return
	}

	var req model.CreateWorkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workout, err := h.workoutService.Create(trainerID, &req)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrStudentNotLinked):
			c.JSON(http.StatusForbidden, gin.H{"error": "esse aluno não está vinculado à sua conta"})
		case errors.Is(err, service.ErrExerciseNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": "um ou mais exercícios não foram encontrados"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create workout"})
		}
		return
	}

	c.JSON(http.StatusCreated, workout)
}

func (h *WorkoutHandler) List(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	if role == string(model.RoleTrainer) {
		var studentID *uuid.UUID
		if raw := c.Query("student_id"); raw != "" {
			id, err := uuid.Parse(raw)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student_id"})
				return
			}
			studentID = &id
		}

		workouts, err := h.workoutService.ListForTrainer(userID, studentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load workouts"})
			return
		}
		c.JSON(http.StatusOK, workouts)
		return
	}

	workouts, err := h.workoutService.ListForStudent(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load workouts"})
		return
	}
	c.JSON(http.StatusOK, workouts)
}

func (h *WorkoutHandler) Get(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	workout, err := h.workoutService.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrWorkoutNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load workout"})
		return
	}

	if workout.TrainerID != userID && workout.StudentID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	c.JSON(http.StatusOK, workout)
}

func (h *WorkoutHandler) Delete(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.workoutService.Delete(trainerID, id); err != nil {
		switch {
		case errors.Is(err, repository.ErrWorkoutNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
		case errors.Is(err, service.ErrWorkoutForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete workout"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *WorkoutHandler) Complete(c *gin.Context) {
	studentID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.workoutService.Complete(studentID, id); err != nil {
		switch {
		case errors.Is(err, repository.ErrWorkoutNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
		case errors.Is(err, service.ErrWorkoutForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to complete workout"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
