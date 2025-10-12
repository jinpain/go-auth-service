package authhttpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/auth"
	"github.com/jinpain/go-auth-service/internal/session"
	"github.com/jinpain/go-auth-service/internal/user"
)

type AuthService interface {
	Register(user *user.Model) error
	Login(email, password string, session *session.Model) (*auth.Model, error)
}

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var registerDTO RegisterRequest

	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := user.Model{
		Email:    registerDTO.Email,
		Phone:    registerDTO.Phone,
		Password: registerDTO.Password,
	}

	if err := h.authService.Register(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user registered"})
}

func (h *Handler) Login(c *gin.Context) {
	var loginDTO LoginRequest

	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := session.Model{
		Device:    c.Request.UserAgent(),
		IpAddress: c.ClientIP(),
	}

	authTokens, err := h.authService.Login(loginDTO.Email, loginDTO.Password, &session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		AccessToken:  authTokens.AccessToken,
		RefreshToken: authTokens.RefreshToken,
	})
}
