package model

import (
	"time"

	"github.com/google/uuid"
)

type ProgressPhoto struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	StudentID    uuid.UUID `gorm:"type:uuid;not null;index" json:"student_id"`
	UploadedByID uuid.UUID `gorm:"type:uuid;not null" json:"uploaded_by_id"`
	ImagePath    string    `gorm:"not null" json:"image_path"`
	RecordedAt   time.Time `gorm:"not null" json:"recorded_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (ProgressPhoto) TableName() string {
	return "progress_photos"
}
