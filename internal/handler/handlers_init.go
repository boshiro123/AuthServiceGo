package handler

import (
	model "auth-service-go/internal/models"
	"auth-service-go/pkg/config"
	"context"
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type Authorization interface {
	Register(ctx context.Context, user model.RegisterRequest) (model.Tokens, error)
	Login(ctx context.Context, user model.LoginRequest) (model.Tokens, error)
	Ping(ctx context.Context) (string, error)
	Refresh(ctx context.Context, user model.RefreshRequest) (model.Tokens, error)
	// SignIn(ctx context.Context, user model.SignInRequest) (model.Tokens, error)
}

type Handler struct {
	authorization Authorization
	logger        *slog.Logger
}

func NewHandler(authorization Authorization, log *slog.Logger) *Handler {
	return &Handler{authorization: authorization, logger: log}
}

func (h *Handler) InitRoutes() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		h.logger.Info("Request", slog.Any("method", c.Request.Method), slog.Any("path", c.Request.URL.Path))
		c.Next()

		h.logger.Info("Response", slog.Any("status", c.Writer.Status()))
	})

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOrigins:     []string{"https://lglhub.com", "https://www.lglhub.com", "http://localhost:4200"},
	}))

	api := router.Group("/api")

	// Публичные роуты
	api.POST("/register", h.registerHandler)
	api.POST("/ping", h.pingHandler)
	api.POST("/login", h.loginHandler)
	api.POST("/refresh", h.refreshHandler)

	// Защищённые роуты (требуют JWT-аутентификации)
	// protected := api.Group("/")
	// protected.Use(middleware.JWTAuthMiddleware())
	// protected.POST("/upload_doc", h.uploadDocumentHandler)
	// protected.GET("/get_profile/:user_id", h.getProfileHandler)
	// protected.POST("/update_document_privacy/:document_id", h.updateDocumentPrivacy)
	// protected.POST("/update_document/:document_id", h.updateDocumentHandler)
	// protected.POST("/delete_document/:document_id", h.deleteDocumentHandler)
	// protected.POST("/upload_avatar", h.uploadAvatarHandler)
	// protected.POST("/upload_pdf", h.uploadPDFHandler)
	// protected.GET("/get_doc/:id/download", h.downloadDocumentHandler)

	// Swagger
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// h.logger.Info("Routes initialized", slog.String("func", "InitRoutes"))

	// Получаем порт из конфигурации
	cfg := config.MustLoad()
	router.Run(":" + cfg.App.Port)
}
