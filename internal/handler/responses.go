package handler

import (
	model "auth-service-go/internal/models"
)

type ErrorResponse struct {
	Error          string `json:"error"`
	RequiredAmount int    `json:"required_amount,omitempty"`
}

type RegisterResponse struct {
	Tokens model.Tokens `json:"tokens"`
}

type LoginResponse struct {
	Tokens model.Tokens `json:"tokens"`
}

type RefreshResponse struct {
	Tokens model.Tokens `json:"tokens"`
}

// type SignInResponse struct {
// 	Tokens model.Tokens `json:"tokens"`
// }

type SuccessResponse struct {
	Message string
}
