package service

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

var (
	ErrInvalidImageType = errors.New("invalid image type")
	ErrImageTooLarge    = errors.New("image exceeds maximum allowed size")
	ErrForbiddenUpload  = errors.New("not allowed to upload photos for this student")
)

const (
	maxImageSize = 5 << 20 // 5MB
	uploadsDir   = "uploads/progress-photos"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

type ProgressPhotoService struct {
	progressPhotoRepo *repository.ProgressPhotoRepository
	trainerStudent    *repository.TrainerStudentRepository
}

func NewProgressPhotoService(progressPhotoRepo *repository.ProgressPhotoRepository, trainerStudentRepo *repository.TrainerStudentRepository) *ProgressPhotoService {
	return &ProgressPhotoService{progressPhotoRepo: progressPhotoRepo, trainerStudent: trainerStudentRepo}
}

func (s *ProgressPhotoService) Upload(uploaderID uuid.UUID, uploaderRole string, studentID uuid.UUID, recordedAt time.Time, file *multipart.FileHeader) (*model.ProgressPhoto, error) {
	if uploaderRole == string(model.RoleTrainer) {
		linked, err := s.trainerStudent.ExistsByTrainerAndStudent(uploaderID, studentID)
		if err != nil {
			return nil, err
		}
		if !linked {
			return nil, ErrForbiddenUpload
		}
	} else if uploaderID != studentID {
		return nil, ErrForbiddenUpload
	}

	if file.Size > maxImageSize {
		return nil, ErrImageTooLarge
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return nil, ErrInvalidImageType
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	sniff := make([]byte, 512)
	n, err := src.Read(sniff)
	if err != nil && err != io.EOF {
		return nil, err
	}
	if !strings.HasPrefix(http.DetectContentType(sniff[:n]), "image/") {
		return nil, ErrInvalidImageType
	}
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		return nil, err
	}

	filename := uuid.New().String() + ext
	destPath := filepath.Join(uploadsDir, filename)

	dst, err := os.Create(destPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	if recordedAt.IsZero() {
		recordedAt = time.Now()
	}

	photo := &model.ProgressPhoto{
		ID:           uuid.New(),
		StudentID:    studentID,
		UploadedByID: uploaderID,
		ImagePath:    fmt.Sprintf("/uploads/progress-photos/%s", filename),
		RecordedAt:   recordedAt,
	}
	if err := s.progressPhotoRepo.Create(photo); err != nil {
		return nil, err
	}
	return photo, nil
}

func (s *ProgressPhotoService) ListForStudent(studentID uuid.UUID) ([]model.ProgressPhoto, error) {
	return s.progressPhotoRepo.ListByStudent(studentID)
}

func (s *ProgressPhotoService) ListForTrainerAndStudent(trainerID, studentID uuid.UUID) ([]model.ProgressPhoto, error) {
	linked, err := s.trainerStudent.ExistsByTrainerAndStudent(trainerID, studentID)
	if err != nil {
		return nil, err
	}
	if !linked {
		return nil, ErrForbiddenUpload
	}
	return s.progressPhotoRepo.ListByStudent(studentID)
}
