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

type TrainerStudentHandler struct {
	trainerStudentService *service.TrainerStudentService
}

func NewTrainerStudentHandler(trainerStudentService *service.TrainerStudentService) *TrainerStudentHandler {
	return &TrainerStudentHandler{trainerStudentService: trainerStudentService}
}

func (h *TrainerStudentHandler) List(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	students, err := h.trainerStudentService.List(trainerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load students"})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *TrainerStudentHandler) Add(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	var req model.AddStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := h.trainerStudentService.AddStudent(trainerID, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrStudentLimitReached):
			c.JSON(http.StatusForbidden, gin.H{"error": "limite de alunos do plano atual atingido, faça upgrade para adicionar mais alunos"})
		case errors.Is(err, service.ErrStudentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "nenhum aluno encontrado com esse e-mail"})
		case errors.Is(err, service.ErrStudentAlreadyLinked):
			c.JSON(http.StatusConflict, gin.H{"error": "esse aluno já está vinculado à sua conta"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add student"})
		}
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *TrainerStudentHandler) MyTrainers(c *gin.Context) {
	studentID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	trainers, err := h.trainerStudentService.ListTrainersForStudent(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load trainers"})
		return
	}

	c.JSON(http.StatusOK, trainers)
}
