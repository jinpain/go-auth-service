package tokenhttpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/token"
)

type TokenService interface {
	RefreshToken(token *token.Model) error
}

type Handler struct {
	tokenService TokenService
}

func NewHandler(tokenService TokenService) *Handler {
	return &Handler{
		tokenService: tokenService,
	}
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := token.Model{
		UserID:       c.GetString("user_id"),
		SessionID:    c.GetString("session_id"),
		RefreshToken: req.RefreshToken,
	}

	if err := h.tokenService.RefreshToken(&token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	resp := RefreshTokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	c.JSON(http.StatusOK, resp)
}
