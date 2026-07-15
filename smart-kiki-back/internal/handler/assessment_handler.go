package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/service"
	"github.com/smartkiki/api/pkg/middleware"
)

type AssessmentHandler struct {
	assessmentService *service.AssessmentService
}

func NewAssessmentHandler(assessmentService *service.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{assessmentService: assessmentService}
}

func (h *AssessmentHandler) Create(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	studentID, err := uuid.Parse(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid studentId"})
		return
	}

	var req model.CreateAssessmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assessment, err := h.assessmentService.Create(trainerID, studentID, &req)
	if err != nil {
		if errors.Is(err, service.ErrStudentNotLinked) {
			c.JSON(http.StatusForbidden, gin.H{"error": "esse aluno não está vinculado à sua conta"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create assessment"})
		return
	}

	c.JSON(http.StatusCreated, assessment)
}

func (h *AssessmentHandler) ListMine(c *gin.Context) {
	studentID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	assessments, err := h.assessmentService.ListForStudent(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load assessments"})
		return
	}

	c.JSON(http.StatusOK, assessments)
}
