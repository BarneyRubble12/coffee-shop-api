package service

import (
	"coffee-shop-api/internal/model"
	"coffee-shop-api/internal/repository"
)

// CoffeeService handles business logic for coffee operations
type CoffeeService struct {
	repo repository.CoffeeRepository
}

// NewCoffeeService creates a new coffee service instance
func NewCoffeeService(repo repository.CoffeeRepository) *CoffeeService {
	return &CoffeeService{repo: repo}
}

// GetAllCoffees retrieves all coffees
func (s *CoffeeService) GetAllCoffees() ([]model.Coffee, error) {
	return s.repo.GetAllCoffees()
}

// GetCoffeeByID retrieves a coffee by its ID
func (s *CoffeeService) GetCoffeeByID(id int64) (*model.Coffee, error) {
	return s.repo.GetCoffeeByID(id)
}

// CreateCoffee creates a new coffee
func (s *CoffeeService) CreateCoffee(coffee *model.Coffee) error {
	return s.repo.CreateCoffee(coffee)
}

// UpdateCoffee updates an existing coffee
func (s *CoffeeService) UpdateCoffee(coffee *model.Coffee) error {
	return s.repo.UpdateCoffee(coffee)
}

// DeleteCoffee deletes a coffee by its ID
func (s *CoffeeService) DeleteCoffee(id int64) error {
	return s.repo.DeleteCoffee(id)
}

// GetRepository returns the repository instance for cleanup
func (s *CoffeeService) GetRepository() repository.CoffeeRepository {
	return s.repo
}
