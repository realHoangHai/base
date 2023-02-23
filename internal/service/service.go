package service

import (
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/authenticator/internal/middleware"
	"github.com/realHoangHai/authenticator/internal/repo"
)

type Service struct {
	r repo.IRepo
	m *middleware.Middleware
}

func NewService(repo repo.IRepo, m *middleware.Middleware) *Service {
	return &Service{
		r: repo,
		m: m,
	}
}

func (s Service) Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "hello")
	}
}
