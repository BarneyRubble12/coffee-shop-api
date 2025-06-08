//go:build wireinject
// +build wireinject

package di

import (
	"coffee-shop-api/internal/handler"
	"coffee-shop-api/internal/repository"
	"coffee-shop-api/internal/service"

	"github.com/google/wire"
)

func InitializeAPI() (*handler.CoffeeHandler, error) {
	wire.Build(
		provideDBPath,
		repository.NewSQLiteCoffeeRepository,
		service.NewCoffeeService,
		handler.NewCoffeeHandler,
		wire.Bind(new(repository.CoffeeRepository), new(*repository.SQLiteCoffeeRepository)),
	)
	return nil, nil
}
