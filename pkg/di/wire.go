//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"my-clean.go/pkg/api"
	"my-clean.go/pkg/api/handler"
	"my-clean.go/pkg/config"
	"my-clean.go/pkg/db"
	"my-clean.go/pkg/repository"
	"my-clean.go/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,

		 repository.NewUserRepository,

		 usecase.NewUserUseCase,

		 handler.NewUserHandler,

		 http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
