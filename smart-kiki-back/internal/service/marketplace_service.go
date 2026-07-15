package service

import (
	"sort"

	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
)

type MarketplaceService struct {
	subscriptionRepo *repository.SubscriptionRepository
	userRepo         *repository.UserRepository
}

func NewMarketplaceService(subscriptionRepo *repository.SubscriptionRepository, userRepo *repository.UserRepository) *MarketplaceService {
	return &MarketplaceService{subscriptionRepo: subscriptionRepo, userRepo: userRepo}
}

func (s *MarketplaceService) ListVisibleTrainers() ([]model.MarketplaceTrainer, error) {
	subs, err := s.subscriptionRepo.ListAll()
	if err != nil {
		return nil, err
	}

	trainers := make([]model.MarketplaceTrainer, 0, len(subs))
	for _, sub := range subs {
		if !model.IsVisibleInMarketplace(sub) {
			continue
		}
		user, err := s.userRepo.FindByID(sub.TrainerID)
		if err != nil {
			continue
		}
		trainers = append(trainers, model.MarketplaceTrainer{User: *user, Plan: sub.Plan})
	}

	sort.SliceStable(trainers, func(i, j int) bool {
		return trainers[i].Plan == model.PlanPro && trainers[j].Plan != model.PlanPro
	})

	return trainers, nil
}
