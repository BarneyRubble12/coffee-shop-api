package di

import (
	"coffee-shop-api/internal/handler"
	"coffee-shop-api/internal/repository"
	"coffee-shop-api/internal/service"

	"github.com/google/wire"
)

// provideAPI creates a new API instance with all dependencies wired
var provideAPI = wire.NewSet(
	repository.NewInMemoryCoffeeRepository,
	wire.Bind(new(repository.CoffeeRepository), new(*repository.InMemoryCoffeeRepository)),
	service.NewCoffeeService,
	handler.NewCoffeeHandler,
)
