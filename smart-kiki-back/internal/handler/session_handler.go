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

type SessionHandler struct {
	sessionService *service.SessionService
}

func NewSessionHandler(sessionService *service.SessionService) *SessionHandler {
	return &SessionHandler{sessionService: sessionService}
}

func (h *SessionHandler) Create(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	if role, _ := c.Get(middleware.ContextRoleKey); role != string(model.RoleTrainer) {
		c.JSON(http.StatusForbidden, gin.H{"error": "only trainers can schedule sessions"})
		return
	}

	var req model.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.sessionService.Create(trainerID, &req)
	if err != nil {
		if errors.Is(err, service.ErrStudentNotLinked) {
			c.JSON(http.StatusForbidden, gin.H{"error": "esse aluno não está vinculado à sua conta"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	c.JSON(http.StatusCreated, session)
}

func (h *SessionHandler) List(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	var (
		sessions []model.Session
		err      error
	)
	if role == string(model.RoleTrainer) {
		sessions, err = h.sessionService.ListForTrainer(userID)
	} else {
		sessions, err = h.sessionService.ListForStudent(userID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load sessions"})
		return
	}

	c.JSON(http.StatusOK, sessions)
}

func (h *SessionHandler) Cancel(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.sessionService.Cancel(userID, id); err != nil {
		switch {
		case errors.Is(err, repository.ErrSessionNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		case errors.Is(err, service.ErrSessionForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cancel session"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
