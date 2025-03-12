package services

import (
	"auth-service-go/internal/grpc/auth"
	model "auth-service-go/internal/models"
	pkgauth "auth-service-go/pkg/auth"
	"context"

	"github.com/google/uuid"
)

// Убедимся, что AuthService реализует интерфейс auth.Auth
var _ auth.Auth = (*AuthService)(nil)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
}

type AuthService struct {
	authorization Authorization
}

func NewAuthService(authorization Authorization) *AuthService {
	return &AuthService{authorization: authorization}
}

// Реализация интерфейса auth.Auth
func (s *AuthService) Login(ctx context.Context, email string, password string, appID int) (token string, err error) {
	// TODO: Implement login
	return "", nil
}

func (s *AuthService) RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error) {
	// TODO: Implement register
	return 0, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, userID int64) (isAdmin bool, err error) {
	// TODO: Implement isAdmin check
	return false, nil
}

// Существующий метод
func (s *AuthService) RegisterUser(ctx context.Context, user model.User) error {
	userID := uuid.New()

	u := model.User{
		ID:           userID,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: pkgauth.GeneratePasswordHash(user.PasswordHash),
	}

	return s.authorization.CreateUser(ctx, u)
}
