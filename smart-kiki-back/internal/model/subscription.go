package model

import (
	"time"

	"github.com/google/uuid"
)

type PlanType string

const (
	PlanFree PlanType = "free"
	PlanPlus PlanType = "plus"
	PlanPro  PlanType = "pro"
)

const plusVisibilityWindow = 7 * 24 * time.Hour

type TrainerSubscription struct {
	TrainerID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"trainer_id"`
	Plan               PlanType  `gorm:"type:varchar(10);not null;default:free" json:"plan"`
	CurrentPeriodStart time.Time `json:"current_period_start"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (TrainerSubscription) TableName() string {
	return "trainer_subscriptions"
}

// StudentLimit retorna o limite de alunos do plano, ou nil para ilimitado.
func StudentLimit(plan PlanType) *int {
	switch plan {
	case PlanFree:
		limit := 2
		return &limit
	case PlanPlus:
		limit := 10
		return &limit
	default:
		return nil
	}
}

func PlanPrice(plan PlanType) float64 {
	switch plan {
	case PlanPlus:
		return 10
	case PlanPro:
		return 20
	default:
		return 0
	}
}

// IsVisibleInMarketplace decide se o personal aparece na busca com base no plano
// e, no caso do plano Plus, na janela de 7 dias a partir do início do ciclo atual.
func IsVisibleInMarketplace(sub TrainerSubscription) bool {
	switch sub.Plan {
	case PlanPro:
		return true
	case PlanPlus:
		return time.Since(sub.CurrentPeriodStart) <= plusVisibilityWindow
	default:
		return false
	}
}

func IsValidPlan(plan PlanType) bool {
	switch plan {
	case PlanFree, PlanPlus, PlanPro:
		return true
	default:
		return false
	}
}

type ChangePlanRequest struct {
	Plan PlanType `json:"plan" binding:"required"`
}

type SubscriptionResponse struct {
	Plan                 string    `json:"plan"`
	Price                float64   `json:"price"`
	StudentLimit         *int      `json:"student_limit"`
	StudentsCount        int64     `json:"students_count"`
	VisibleInMarketplace bool      `json:"visible_in_marketplace"`
	CurrentPeriodStart   time.Time `json:"current_period_start"`
}

type MarketplaceTrainer struct {
	User
	Plan PlanType `json:"plan"`
}
