package main

import (
	"log"
	"os"

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
	exerciseRepo := repository.NewExerciseRepository(db)
	workoutRepo := repository.NewWorkoutRepository(db)
	workoutLogRepo := repository.NewWorkoutLogRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	nutritionPlanRepo := repository.NewNutritionPlanRepository(db)
	assessmentRepo := repository.NewAssessmentRepository(db)
	progressPhotoRepo := repository.NewProgressPhotoRepository(db)

	authService := service.NewAuthService(userRepo, subscriptionRepo, cfg.JWT.Secret, cfg.JWT.Expiration)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo, trainerStudentRepo)
	trainerStudentService := service.NewTrainerStudentService(trainerStudentRepo, subscriptionRepo, userRepo)
	exerciseService := service.NewExerciseService(exerciseRepo)
	workoutService := service.NewWorkoutService(workoutRepo, workoutLogRepo, exerciseRepo, trainerStudentRepo)
	sessionService := service.NewSessionService(sessionRepo, trainerStudentRepo)
	messageService := service.NewMessageService(messageRepo, trainerStudentRepo)
	marketplaceService := service.NewMarketplaceService(subscriptionRepo, userRepo)
	nutritionPlanService := service.NewNutritionPlanService(nutritionPlanRepo, trainerStudentRepo)
	assessmentService := service.NewAssessmentService(assessmentRepo, trainerStudentRepo)
	progressPhotoService := service.NewProgressPhotoService(progressPhotoRepo, trainerStudentRepo)

	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userRepo)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)
	trainerStudentHandler := handler.NewTrainerStudentHandler(trainerStudentService)
	exerciseHandler := handler.NewExerciseHandler(exerciseService)
	workoutHandler := handler.NewWorkoutHandler(workoutService)
	sessionHandler := handler.NewSessionHandler(sessionService)
	messageHandler := handler.NewMessageHandler(messageService)
	marketplaceHandler := handler.NewMarketplaceHandler(marketplaceService, trainerStudentService)
	nutritionPlanHandler := handler.NewNutritionPlanHandler(nutritionPlanService)
	assessmentHandler := handler.NewAssessmentHandler(assessmentService)
	progressPhotoHandler := handler.NewProgressPhotoHandler(progressPhotoService)

	if err := os.MkdirAll("uploads/progress-photos", 0o755); err != nil {
		log.Fatalf("failed to create uploads directory: %v", err)
	}

	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(middleware.CORS(cfg.CORSOrigin))

	router.GET("/health", handler.Health)
	router.Static("/uploads", "./uploads")

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
			trainer.PUT("/students/:studentId/nutrition-plan", nutritionPlanHandler.Upsert)
			trainer.GET("/students/:studentId/nutrition-plan", nutritionPlanHandler.GetForStudent)
			trainer.POST("/students/:studentId/assessments", assessmentHandler.Create)
		}

		students := api.Group("/students")
		students.Use(middleware.Auth(cfg.JWT.Secret))
		{
			students.GET("/me/trainers", trainerStudentHandler.MyTrainers)
			students.GET("/me/nutrition-plan", nutritionPlanHandler.GetMine)
			students.GET("/me/assessments", assessmentHandler.ListMine)
		}

		exercises := api.Group("/exercises")
		exercises.Use(middleware.Auth(cfg.JWT.Secret))
		{
			exercises.GET("", exerciseHandler.List)
		}

		workouts := api.Group("/workouts")
		workouts.Use(middleware.Auth(cfg.JWT.Secret))
		{
			workouts.POST("", workoutHandler.Create)
			workouts.GET("", workoutHandler.List)
			workouts.GET("/:id", workoutHandler.Get)
			workouts.DELETE("/:id", workoutHandler.Delete)
			workouts.POST("/:id/complete", workoutHandler.Complete)
		}

		sessions := api.Group("/sessions")
		sessions.Use(middleware.Auth(cfg.JWT.Secret))
		{
			sessions.POST("", sessionHandler.Create)
			sessions.GET("", sessionHandler.List)
			sessions.POST("/:id/cancel", sessionHandler.Cancel)
		}

		messages := api.Group("/messages")
		messages.Use(middleware.Auth(cfg.JWT.Secret))
		{
			messages.POST("", messageHandler.Send)
			messages.GET("", messageHandler.ListConversation)
		}

		marketplace := api.Group("/marketplace")
		marketplace.Use(middleware.Auth(cfg.JWT.Secret))
		{
			marketplace.GET("/trainers", marketplaceHandler.List)
			marketplace.POST("/trainers/:id/request", middleware.RequireRole("client"), marketplaceHandler.Request)
		}

		progressPhotos := api.Group("/progress-photos")
		progressPhotos.Use(middleware.Auth(cfg.JWT.Secret))
		{
			progressPhotos.POST("", progressPhotoHandler.Upload)
			progressPhotos.GET("", progressPhotoHandler.List)
		}
	}

	log.Printf("smart-kiki-api listening on :%s (%s)", cfg.Server.Port, cfg.Server.Env)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
