package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realHoangHai/authenticator/config"
	"github.com/realHoangHai/authenticator/internal/middleware"
	"github.com/realHoangHai/authenticator/internal/service"
	"github.com/realHoangHai/authenticator/pkg/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type Server struct {
	addr string
	m    *middleware.Middleware
	s    *service.Service
}

func NewServer(m *middleware.Middleware, s *service.Service) *Server {
	addr := fmt.Sprintf(":%d", config.C.AppConfig.Port)
	return &Server{
		addr: addr,
		m:    m,
		s:    s,
	}
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	return s.start(ctx, srv)
}

func (s *Server) start(ctx context.Context, srv *http.Server) error {
	var g errgroup.Group

	g.Go(func() error {
		<-ctx.Done()
		timeout := time.Duration(config.C.AppConfig.ShutdownTimeout) * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		return srv.Shutdown(ctx)
	})

	g.Go(func() error {
		log.I("Starting server on http://localhost%s", s.addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	return g.Wait()
}

func (s *Server) router() http.Handler {
	gin.SetMode(config.C.AppConfig.Mode)
	r := gin.New()
	r.NoMethod(s.m.NoMethod())
	r.NoRoute(s.m.NoRoute())
	r.Use(s.m.Recover())
	r.Use(s.m.Cors())

	swaggerURL := ginSwagger.URL(fmt.Sprintf("0.0.0.0%s/swagger/doc.json", s.addr)) // the  url poiting to API definition
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// register api
	g := r.Group("/api/")

	v1 := g.Group("/v1/")
	{
		v1.POST("/login", s.s.Login())
	}

	return r
}
