package service

import (
	"coffee-shop-api/internal/model"
	"coffee-shop-api/internal/repository"
)

// CoffeeService handles business logic for coffee operations
type CoffeeService struct {
	repo repository.CoffeeRepository
}

// NewCoffeeService creates a new instance of CoffeeService
func NewCoffeeService(repo repository.CoffeeRepository) *CoffeeService {
	return &CoffeeService{
		repo: repo,
	}
}

// GetAllCoffees returns all coffees
func (s *CoffeeService) GetAllCoffees() []model.Coffee {
	return s.repo.GetAll()
}

// GetCoffeeByID returns a coffee by its ID
func (s *CoffeeService) GetCoffeeByID(id int) (model.Coffee, error) {
	return s.repo.GetByID(id)
}

// CreateCoffee creates a new coffee
func (s *CoffeeService) CreateCoffee(coffee model.Coffee) model.Coffee {
	return s.repo.Create(coffee)
}

// UpdateCoffee updates an existing coffee
func (s *CoffeeService) UpdateCoffee(id int, coffee model.Coffee) (model.Coffee, error) {
	return s.repo.Update(id, coffee)
}

// DeleteCoffee deletes a coffee by its ID
func (s *CoffeeService) DeleteCoffee(id int) error {
	return s.repo.Delete(id)
}
