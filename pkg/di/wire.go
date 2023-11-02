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

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabse,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewSupAdminRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewSupAdminUseCase,
		usecase.NewProductUsecase,
		usecase.NewCartUsecase,
		usecase.NewOrderUseCase,
		usecase.NewPaymentuseCase,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewSupAdminHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		http.NewServerHTTP,
	)
	return &http.ServerHTTP{}, nil
}
