package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Ping
// @Description Проверка доступности сервиса
// @Tags service
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [post]
func (h *Handler) pingHandler(c *gin.Context) {
	// answer, err := h.authorization.Ping(c.Request.Context())
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}
