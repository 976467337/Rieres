package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartkiki/api/internal/service"
)

type ExerciseHandler struct {
	exerciseService *service.ExerciseService
}

func NewExerciseHandler(exerciseService *service.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{exerciseService: exerciseService}
}

func (h *ExerciseHandler) List(c *gin.Context) {
	muscleGroup := c.Query("muscle_group")

	exercises, err := h.exerciseService.List(muscleGroup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load exercises"})
		return
	}

	c.JSON(http.StatusOK, exercises)
}
