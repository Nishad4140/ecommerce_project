// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Nishad4140/ecommerce_project/pkg/api"
	"github.com/Nishad4140/ecommerce_project/pkg/api/handler"
	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/Nishad4140/ecommerce_project/pkg/db"
	"github.com/Nishad4140/ecommerce_project/pkg/repository"
	"github.com/Nishad4140/ecommerce_project/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI1(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabse(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUsecase := usecase.NewCartUsecase(cartRepository)
	userHandler := handler.NewUserHandler(userUseCase, cartUsecase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	productRepository := repository.NewProductRepository(gormDB)
	productUsecase := usecase.NewProductUsecase(productRepository)
	productHandler := handler.NewProductHandler(productUsecase)
	cartHandler := handler.NewCartHandler(cartUsecase)
	supAdminRepository := repository.NewSupAdminRepository(gormDB)
	supAdminUseCase := usecase.NewSupAdminUseCase(supAdminRepository)
	supAdminHandler := handler.NewSupAdminHandler(supAdminUseCase)
	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentuseCase(paymentRepository, orderRepository, cfg)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, productHandler, cartHandler, supAdminHandler, orderHandler, paymentHandler)
	return serverHTTP, nil
}
