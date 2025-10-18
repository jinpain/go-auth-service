package sessionhttpv1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/session"
)

func SetupRouter(r *gin.RouterGroup, sessionService *session.Service) {
	h := NewHandler(sessionService)

	session := r.Group("/v1/session")
	{
		session.POST("/revoke/:session_id", h.RevokeSession)
		session.POST("/logout", h.Logout)
	}
}
