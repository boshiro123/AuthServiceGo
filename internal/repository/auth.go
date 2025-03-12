package repository

import (
	sqlc "auth-service-go/internal/db/sqlc"
	model "auth-service-go/internal/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db    *pgxpool.Pool
	query *sqlc.Queries
}

func NewAuth(db *pgxpool.Pool, query *sqlc.Queries) *AuthRepository {
	return &AuthRepository{
		db:    db,
		query: query,
	}
}

func (a *AuthRepository) CreateUser(ctx context.Context, user model.User) error {
	err := a.query.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       uuid.New(),
		Email:    user.Email,
		Name:     user.Name,
		Password: user.PasswordHash,
	})
	if err != nil {
		return err
	}
	return nil
}
