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

type MessageHandler struct {
	messageService *service.MessageService
}

func NewMessageHandler(messageService *service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) Send(c *gin.Context) {
	senderID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	var req model.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, err := h.messageService.Send(senderID, role.(string), req.ToUserID, req.Body)
	if err != nil {
		if errors.Is(err, service.ErrStudentNotLinked) {
			c.JSON(http.StatusForbidden, gin.H{"error": "essa conversa não está disponível"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) ListConversation(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	otherUserID, err := uuid.Parse(c.Query("with"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'with' query param"})
		return
	}

	messages, err := h.messageService.ListConversation(userID, role.(string), otherUserID)
	if err != nil {
		if errors.Is(err, service.ErrStudentNotLinked) {
			c.JSON(http.StatusForbidden, gin.H{"error": "essa conversa não está disponível"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load conversation"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
