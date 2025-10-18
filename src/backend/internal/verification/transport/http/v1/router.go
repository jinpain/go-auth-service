package verificationhttpv1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinpain/go-auth-service/internal/verification"
)

func SetupRouter(r *gin.RouterGroup, s *verification.Service) {
	h := NewHandler(s)

	verification := r.Group("/v1/verification")
	{
		verification.POST("/check/:code", h.CheckVerifyCode)
	}
}
