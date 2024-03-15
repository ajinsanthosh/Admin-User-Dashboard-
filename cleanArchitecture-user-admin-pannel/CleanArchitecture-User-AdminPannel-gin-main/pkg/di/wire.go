//go:build wireinject
// +build wireinject

package di

import (
	http "example/pkg/api"
	"example/pkg/api/handler"
	config "example/pkg/config"
	"example/pkg/db"
	"example/pkg/repository"
	"example/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserDataBase,
		repository.NewAdminRepository,

		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,

		handler.NewUserHandler,
		handler.NewAdminHandler,

		http.NewServerHttp)

	return &http.ServerHTTP{}, nil
}
