package middleware

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/authenticator/internal/repo"
	"github.com/realHoangHai/authenticator/pkg/errors"
)

type Middleware struct {
	ctx context.Context
	r   repo.IRepo
}

func NewMiddleware(ctx context.Context, r repo.IRepo) *Middleware {
	return &Middleware{
		ctx: ctx,
		r:   r,
	}
}

func (m *Middleware) Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*errors.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}

				appErr := errors.ErrInternal(fmt.Errorf("%v", err))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				return
			}
		}()

		c.Next()
	}
}

func (m *Middleware) Empty() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (m *Middleware) Cors() gin.HandlerFunc {
	return cors.Default()
}

func (m *Middleware) NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(errors.ErrEntityNotFound("page", fmt.Errorf("not found")))
	}

}

func (m *Middleware) NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic(errors.ErrMethodNotAllowed)
	}
}
