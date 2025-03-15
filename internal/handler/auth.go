package handler

import (
	model "auth-service-go/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Sign Up
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.SignUpRequest true "User Sign Up Data"
// @Success 200 {object} SignUpResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sign_up [post]
func (h *Handler) registerHandler(c *gin.Context) {
	var user model.RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	tokens, err := h.authorization.Register(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{Tokens: tokens})
}

func (h *Handler) loginHandler(c *gin.Context) {
	var user model.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	tokens, err := h.authorization.Login(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Tokens: tokens})
}

func (h *Handler) refreshHandler(c *gin.Context) {
	var user model.RefreshRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	tokens, err := h.authorization.Refresh(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RefreshResponse{Tokens: tokens})
}

// @Summary Sign In
// @Description Authenticate a user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.SignInRequest true "User Sign In Data"
// @Success 200 {object} SignInResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /sign_in [post]
// func (h *Handler) signInHandler(c *gin.Context) {
// 	var user model.SignInRequest
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	tokens, err := h.authorization.SignIn(c.Request.Context(), user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, SignInResponse{Tokens: tokens})
// }
