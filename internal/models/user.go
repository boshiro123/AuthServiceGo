package model

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Email        string    `json:"email,omitempty"`
	Name         string    `json:"name,omitempty"`
	Password 		 string    `json:"password,omitempty"`
	DeletedAt    int64     `json:"deleted_at,omitempty"`
	CreatedAt    int64     `json:"created_at,omitempty"`
	UpdatedAt    int64     `json:"updated_at,omitempty"`
}
