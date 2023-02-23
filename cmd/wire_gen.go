// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/realHoangHai/authenticator/internal/middleware"
	"github.com/realHoangHai/authenticator/internal/repo"
	"github.com/realHoangHai/authenticator/internal/server"
	"github.com/realHoangHai/authenticator/internal/service"
)

// Injectors from wire.go:

func initServer(ctx context.Context) (*server.Server, error) {
	iRepo, err := repo.NewRepo(ctx)
	if err != nil {
		return nil, err
	}
	middlewareMiddleware := middleware.NewMiddleware(ctx, iRepo)
	serviceService := service.NewService(iRepo, middlewareMiddleware)
	serverServer := server.NewServer(middlewareMiddleware, serviceService)
	return serverServer, nil
}
