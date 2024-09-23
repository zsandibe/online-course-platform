package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/online-course-platform/internal/domain"
)

func (h *Handler) signIn(c *gin.Context) {
	var inp domain.SignInRequest

	if err := c.ShouldBindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input: %v", err))
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var inp domain.SignUpRequest

	if err := c.ShouldBindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid input: %v", err))
		return
	}
}

func (h *Handler) forgotPassword(c *gin.Context) {
}

func (h *Handler) verifyEmail(c *gin.Context) {
}

func (h *Handler) changePassword(c *gin.Context) {
}

func (h *Handler) updateProfile(c *gin.Context) {
}

func (h *Handler) deleteAccount(c *gin.Context) {}
