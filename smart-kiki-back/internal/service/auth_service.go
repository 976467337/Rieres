package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/smartkiki/api/internal/model"
	"github.com/smartkiki/api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailAlreadyInUse  = errors.New("email already in use")
)

type AuthService struct {
	userRepo         *repository.UserRepository
	subscriptionRepo *repository.SubscriptionRepository
	jwtSecret        string
	jwtExpires       time.Duration
}

func NewAuthService(
	userRepo *repository.UserRepository,
	subscriptionRepo *repository.SubscriptionRepository,
	jwtSecret string,
	jwtExpires time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		subscriptionRepo: subscriptionRepo,
		jwtSecret:        jwtSecret,
		jwtExpires:       jwtExpires,
	}
}

func (s *AuthService) Register(req *model.RegisterRequest) (*model.AuthResponse, error) {
	if _, err := s.userRepo.FindByEmail(req.Email); err == nil {
		return nil, ErrEmailAlreadyInUse
	} else if !errors.Is(err, repository.ErrUserNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:           uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         req.Role,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	if user.Role == model.RoleTrainer {
		if err := s.subscriptionRepo.Create(&model.TrainerSubscription{
			TrainerID:          user.ID,
			Plan:               model.PlanFree,
			CurrentPeriodStart: time.Now(),
		}); err != nil {
			return nil, err
		}
	}

	token, err := s.issueToken(user)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.issueToken(user)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) issueToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID.String(),
		"role": string(user.Role),
		"exp":  time.Now().Add(s.jwtExpires).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
