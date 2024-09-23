package v1

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, errorResponse{Code: code, Message: err.Error()})
}
