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

type NutritionPlanHandler struct {
	nutritionPlanService *service.NutritionPlanService
}

func NewNutritionPlanHandler(nutritionPlanService *service.NutritionPlanService) *NutritionPlanHandler {
	return &NutritionPlanHandler{nutritionPlanService: nutritionPlanService}
}

func (h *NutritionPlanHandler) Upsert(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	studentID, err := uuid.Parse(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid studentId"})
		return
	}

	var req model.UpsertNutritionPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan, err := h.nutritionPlanService.Upsert(trainerID, studentID, req.Content)
	if err != nil {
		if errors.Is(err, service.ErrStudentNotLinked) {
			c.JSON(http.StatusForbidden, gin.H{"error": "esse aluno não está vinculado à sua conta"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save nutrition plan"})
		return
	}

	c.JSON(http.StatusOK, plan)
}

func (h *NutritionPlanHandler) GetForStudent(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	studentID, err := uuid.Parse(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid studentId"})
		return
	}

	plan, err := h.nutritionPlanService.GetForTrainerAndStudent(trainerID, studentID)
	if err != nil {
		if errors.Is(err, repository.ErrNutritionPlanNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "nutrition plan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load nutrition plan"})
		return
	}

	c.JSON(http.StatusOK, plan)
}

func (h *NutritionPlanHandler) GetMine(c *gin.Context) {
	studentID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	plan, err := h.nutritionPlanService.GetLatestForStudent(studentID)
	if err != nil {
		if errors.Is(err, repository.ErrNutritionPlanNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "nutrition plan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load nutrition plan"})
		return
	}

	c.JSON(http.StatusOK, plan)
}
