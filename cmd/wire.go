//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/realHoangHai/authenticator/internal/middleware"
	"github.com/realHoangHai/authenticator/internal/repo"
	"github.com/realHoangHai/authenticator/internal/server"
	"github.com/realHoangHai/authenticator/internal/service"
)

func initServer(ctx context.Context) (*server.Server, error) {
	wire.Build(
		repo.ProviderRepoSet,
		middleware.NewMiddleware,
		service.NewService,
		server.NewServer,
	)

	return new(server.Server), nil
}
