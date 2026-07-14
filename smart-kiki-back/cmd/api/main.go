package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/smartkiki/api/internal/config"
	"github.com/smartkiki/api/internal/handler"
	"github.com/smartkiki/api/internal/repository"
	"github.com/smartkiki/api/internal/service"
	"github.com/smartkiki/api/pkg/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	trainerStudentRepo := repository.NewTrainerStudentRepository(db)

	authService := service.NewAuthService(userRepo, subscriptionRepo, cfg.JWT.Secret, cfg.JWT.Expiration)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo, trainerStudentRepo)
	trainerStudentService := service.NewTrainerStudentService(trainerStudentRepo, subscriptionRepo, userRepo)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userRepo)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)
	trainerStudentHandler := handler.NewTrainerStudentHandler(trainerStudentService)

	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.CORS(cfg.CORSOrigin))

	router.GET("/health", handler.Health)

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		users := api.Group("/users")
		users.Use(middleware.Auth(cfg.JWT.Secret))
		{
			users.GET("/me", userHandler.Me)
		}

		trainer := api.Group("/trainer")
		trainer.Use(middleware.Auth(cfg.JWT.Secret), middleware.RequireRole("trainer"))
		{
			trainer.GET("/subscription", subscriptionHandler.Get)
			trainer.PATCH("/subscription", subscriptionHandler.ChangePlan)
			trainer.GET("/students", trainerStudentHandler.List)
			trainer.POST("/students", trainerStudentHandler.Add)
		}
	}

	log.Printf("smart-kiki-api listening on :%s (%s)", cfg.Server.Port, cfg.Server.Env)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
