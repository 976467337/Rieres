package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/service"
	"github.com/smartkiki/api/pkg/middleware"
)

type ProgressPhotoHandler struct {
	progressPhotoService *service.ProgressPhotoService
}

func NewProgressPhotoHandler(progressPhotoService *service.ProgressPhotoService) *ProgressPhotoHandler {
	return &ProgressPhotoHandler{progressPhotoService: progressPhotoService}
}

func (h *ProgressPhotoHandler) Upload(c *gin.Context) {
	uploaderID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	studentID, err := uuid.Parse(c.PostForm("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student_id"})
		return
	}

	var recordedAt time.Time
	if raw := c.PostForm("recorded_at"); raw != "" {
		recordedAt, err = time.Parse(time.RFC3339, raw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recorded_at"})
			return
		}
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing photo file"})
		return
	}

	photo, err := h.progressPhotoService.Upload(uploaderID, role.(string), studentID, recordedAt, file)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrForbiddenUpload):
			c.JSON(http.StatusForbidden, gin.H{"error": "você não pode enviar fotos para esse aluno"})
		case errors.Is(err, service.ErrInvalidImageType):
			c.JSON(http.StatusBadRequest, gin.H{"error": "arquivo precisa ser uma imagem (jpg, png ou webp)"})
		case errors.Is(err, service.ErrImageTooLarge):
			c.JSON(http.StatusBadRequest, gin.H{"error": "imagem excede o tamanho máximo de 5MB"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload photo"})
		}
		return
	}

	c.JSON(http.StatusCreated, photo)
}

func (h *ProgressPhotoHandler) List(c *gin.Context) {
	userID := c.MustGet(middleware.ContextUserIDKey).(uuid.UUID)
	role, _ := c.Get(middleware.ContextRoleKey)

	if role == string(model.RoleTrainer) {
		raw := c.Query("student_id")
		if raw == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "student_id is required"})
			return
		}
		studentID, err := uuid.Parse(raw)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student_id"})
			return
		}

		photos, err := h.progressPhotoService.ListForTrainerAndStudent(userID, studentID)
		if err != nil {
			if errors.Is(err, service.ErrForbiddenUpload) {
				c.JSON(http.StatusForbidden, gin.H{"error": "esse aluno não está vinculado à sua conta"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load photos"})
			return
		}
		c.JSON(http.StatusOK, photos)
		return
	}

	photos, err := h.progressPhotoService.ListForStudent(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}
