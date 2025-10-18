package authhttpv1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/auth"
)

func SetupRouter(r *gin.RouterGroup, authService *auth.Service) {
	h := NewHandler(authService)

	auth := r.Group("/v1/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
