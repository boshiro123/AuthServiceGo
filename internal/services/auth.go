package services

import (
	model "auth-service-go/internal/models"
	"auth-service-go/pkg/auth"
	"auth-service-go/pkg/config"
	"auth-service-go/pkg/logger"
	"context"
	"log/slog"
	"time"

	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type AuthService struct {
	authorization Authorization
	config        *config.Config
}

func NewAuthService(authorization Authorization) *AuthService {
	return &AuthService{
		authorization: authorization,
		config:        config.MustLoad(),
	}
}

func (s *AuthService) Register(ctx context.Context, user model.RegisterRequest) (model.Tokens, error) {
	userID := uuid.New()

	u := model.User{
		ID:           userID,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: auth.GeneratePasswordHash(user.Password),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	}

	if err := s.authorization.CreateUser(ctx, u); err != nil {
		return model.Tokens{}, err
	}

	accessToken, refreshToken, err := auth.GenerateToken(userID)
	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, user model.LoginRequest) (model.Tokens, error) {
	u, err := s.authorization.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return model.Tokens{}, err
	}

	if !auth.ComparePassword(user.Password, u.PasswordHash) {
		return model.Tokens{}, model.ErrInvalidPassword
	}

	accessToken, refreshToken, err := auth.GenerateToken(u.ID)
	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       u.ID,
	}, nil
}

func (s *AuthService) Refresh(ctx context.Context, user model.RefreshRequest) (model.Tokens, error) {
	log := logger.SetupLogger(s.config.Env)
	claims, err := s.validateRefreshToken(user.RefreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	log.Info("claims", slog.Any("claims", claims))
	// Проверка срока действия (exp claim)
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return model.Tokens{}, errors.New("токен истек")
		}
	}

	// Проверка наличия поля user_id
	userIDValue, ok := claims["user_id"]
	if !ok {
		return model.Tokens{}, errors.New("отсутствует идентификатор пользователя в токене")
	}

	userID, ok := userIDValue.(string)
	if !ok {
		return model.Tokens{}, errors.New("некорректный формат идентификатора пользователя")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return model.Tokens{}, fmt.Errorf("ошибка парсинга UUID: %w", err)
	}

	accessToken, refreshToken, err := auth.GenerateToken(userUUID)
	if err != nil {
		return model.Tokens{}, err
	}

	return model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userUUID,
	}, nil
}

func (s *AuthService) validateRefreshToken(tokenString string) (jwt.MapClaims, error) {
	// Парсинг токена с проверкой подписи
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}

		// Возвращаем секретный ключ для проверки подписи
		return []byte(s.config.Secrets.SecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("ошибка при парсинге токена: %w", err)
	}

	// Проверка валидности токена
	if !token.Valid {
		return nil, errors.New("невалидный токен")
	}

	// Получение claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("невозможно получить claims из токена")
	}

	return claims, nil
}

func (s *AuthService) Ping(ctx context.Context) (string, error) {
	// Здесь можно добавить проверку доступности базы данных или других сервисов
	return "pong", nil
}

// func (s *AuthService) SignIn(ctx context.Context, user model.SignInRequest) (model.Tokens, error) {
// 	u, err := s.authorization.GetUserByEmail(ctx, user.Email)
// 	if err != nil {
// 		return model.Tokens{}, err
// 	}

// 	if !auth.ComparePassword(user.Password, u.PasswordHash) {
// 		return model.Tokens{}, model.ErrInvalidPassword
// 	}

// 	accessToken, refreshToken, err := auth.GenerateToken(u.ID)
// 	if err != nil {
// 		return model.Tokens{}, err
// 	}

// 	return model.Tokens{
// 		AccessToken:  accessToken,
// 		RefreshToken: refreshToken,
// 		UserID:       u.ID,
// 	}, nil
// }
