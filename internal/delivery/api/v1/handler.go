package v1

import (
	"github.com/zsandibe/online-course-platform/internal/service"
	"github.com/zsandibe/online-course-platform/pkg/manager"
)

type Handler struct {
	service      *service.Service
	tokenManager *manager.Manager
}

func NewHandler(service *service.Service, tokenManager *manager.Manager) *Handler {
	return &Handler{service: service, tokenManager: tokenManager}
}
