package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/service"
	"github.com/smartkiki/api/pkg/middleware"
)

type MarketplaceHandler struct {
	marketplaceService    *service.MarketplaceService
	trainerStudentService *service.TrainerStudentService
}

func NewMarketplaceHandler(marketplaceService *service.MarketplaceService, trainerStudentService *service.TrainerStudentService) *MarketplaceHandler {
	return &MarketplaceHandler{marketplaceService: marketplaceService, trainerStudentService: trainerStudentService}
}

func (h *MarketplaceHandler) List(c *gin.Context) {
	trainers, err := h.marketplaceService.ListVisibleTrainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load trainers"})
		return
	}

	c.JSON(http.StatusOK, trainers)
}

func (h *MarketplaceHandler) Request(c *gin.Context) {
	studentID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	trainerID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.trainerStudentService.RequestTrainer(studentID, trainerID); err != nil {
		switch {
		case errors.Is(err, service.ErrStudentLimitReached):
			c.JSON(http.StatusForbidden, gin.H{"error": "esse personal atingiu o limite de alunos do plano atual"})
		case errors.Is(err, service.ErrStudentAlreadyLinked):
			c.JSON(http.StatusConflict, gin.H{"error": "você já está vinculado a esse personal"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to request trainer"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
