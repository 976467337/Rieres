package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"gorm.io/gorm"
)

var ErrSubscriptionNotFound = errors.New("subscription not found")

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(sub *model.TrainerSubscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepository) FindByTrainerID(trainerID uuid.UUID) (*model.TrainerSubscription, error) {
	var sub model.TrainerSubscription
	if err := r.db.Where("trainer_id = ?", trainerID).First(&sub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSubscriptionNotFound
		}
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Update(sub *model.TrainerSubscription) error {
	return r.db.Save(sub).Error
}

func (r *SubscriptionRepository) ListAll() ([]model.TrainerSubscription, error) {
	var subs []model.TrainerSubscription
	err := r.db.Find(&subs).Error
	return subs, err
}
