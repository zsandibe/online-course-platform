package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/online-course-platform/internal/domain"
	validator "github.com/zsandibe/online-course-platform/pkg/validator"
)

func (h *Handler) signIn(c *gin.Context) {
	var inp domain.SignInRequest

	if err := c.ShouldBindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}
}

func (h *Handler) signUp(c *gin.Context) {
	var inp domain.SignUpRequest

	if err := c.ShouldBindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if err := validator.ValidateSignInRequest(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("validation error: %v", err))
		return
	}

	c.JSON(200, &inp)
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
