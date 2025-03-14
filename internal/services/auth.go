package services

import (
	model "auth-service-go/internal/models"
	"auth-service-go/pkg/auth"
	"context"

	"github.com/google/uuid"
)

type Authorization interface {
	CreateUser(ctx context.Context, user model.User) error
	// GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type AuthService struct {
	authorization Authorization
}

func NewAuthService(authorization Authorization) *AuthService {
	return &AuthService{authorization: authorization}
}

func (s *AuthService) Register(ctx context.Context, user model.RegisterRequest) (model.Tokens, error) {
	userID := uuid.New()

	u := model.User{
		ID:           userID,
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: auth.GeneratePasswordHash(user.Password),
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
