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

type SubscriptionHandler struct {
	subscriptionService *service.SubscriptionService
}

func NewSubscriptionHandler(subscriptionService *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{subscriptionService: subscriptionService}
}

func (h *SubscriptionHandler) Get(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	resp, err := h.subscriptionService.Get(trainerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load subscription"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *SubscriptionHandler) ChangePlan(c *gin.Context) {
	trainerID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)

	var req model.ChangePlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.subscriptionService.ChangePlan(trainerID, req.Plan)
	if err != nil {
		if errors.Is(err, service.ErrInvalidPlan) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to change plan"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
