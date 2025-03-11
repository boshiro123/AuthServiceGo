package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
	"context"
)

type AuthRepository struct {
	db *pgxpool.Pool
	query  *sqlc.Queries
}

func NewAuth(db *pgxpool.Pool, query *sqlc.Queries) *AuthRepository {
	return &AuthRepository{
		db: db,
		query:  query,
	}
}

func (a *AuthRepository) CreateUser(ctx context.Context, user model.User) error {
	err := a.query.CreateUser(ctx, sqlc.CreateUserParams{
		ID: pgtype.UUID{Bytes: user.ID, Valid: true},
		Email: user.Email,
		Name: user.Name,
		Password: user.Password,
		CreatedAt: pgtype.Timestamp{Time: time.Now()},
		UpdatedAt: pgtype.Timestamp{Time: time.Now()},
	})
	if err != nil {
		return err
	}
	return nil	
}