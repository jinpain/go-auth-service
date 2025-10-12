package verificationhttpv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerificationService interface {
	CheckVerifyCode(code string) error
}

type Handler struct {
	verificationService VerificationService
}

func NewHandler(verificationService VerificationService) *Handler {
	return &Handler{
		verificationService: verificationService,
	}
}

func (h *Handler) CheckVerifyCode(c *gin.Context) {
	verifyCode := c.Param("code")

	if err := h.verificationService.CheckVerifyCode(verifyCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully verified"})
}
