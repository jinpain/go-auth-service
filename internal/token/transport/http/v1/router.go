package tokenhttpv1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/token"
)

func SetupRouter(r *gin.RouterGroup, tokenService *token.Service) {
	h := NewHandler(tokenService)

	token := r.Group("/v1/token")
	{
		token.POST("/refresh", h.RefreshToken)
	}
}
