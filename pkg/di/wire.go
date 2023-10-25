package di

import (
	http "github.com/Nishad4140/ecommerce_project/pkg/api"
	"github.com/Nishad4140/ecommerce_project/pkg/api/handler"
	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/db"
	"github.com/Nishad4140/ecommerce_project/pkg/repository"
	"github.com/Nishad4140/ecommerce_project/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI1(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabse,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
