package sessionhttpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionService interface {
	RevokeSession(userID, sessionID string) error
}

type Handler struct {
	sessionService SessionService
}

func NewHandler(sessionService SessionService) *Handler {
	return &Handler{
		sessionService: sessionService,
	}
}

func (h *Handler) RevokeSession(c *gin.Context) {
	sessionID := c.Param("session_id")

	userID := c.GetString("user_id")

	if err := h.sessionService.RevokeSession(userID, sessionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "session revoked"})
}

func (h *Handler) Logout(c *gin.Context) {
	sessionID := c.GetString("session_id")

	userID := c.GetString("user_id")

	if err := h.sessionService.RevokeSession(userID, sessionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
