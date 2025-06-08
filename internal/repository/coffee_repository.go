package repository

import (
	"errors"
	"sync"
	"time"

	"coffee-shop-api/internal/model"
)

var (
	ErrCoffeeNotFound = errors.New("coffee not found")
)

// CoffeeRepository defines the interface for coffee data operations
type CoffeeRepository interface {
	GetAllCoffees() ([]model.Coffee, error)
	GetCoffeeByID(id int64) (*model.Coffee, error)
	CreateCoffee(coffee *model.Coffee) error
	UpdateCoffee(coffee *model.Coffee) error
	DeleteCoffee(id int64) error
}

// InMemoryCoffeeRepository implements CoffeeRepository interface using in-memory storage
type InMemoryCoffeeRepository struct {
	coffees map[int64]model.Coffee
	nextID  int64
	mu      sync.RWMutex
}

// NewInMemoryCoffeeRepository creates a new in-memory repository instance
func NewInMemoryCoffeeRepository() *InMemoryCoffeeRepository {
	return &InMemoryCoffeeRepository{
		coffees: make(map[int64]model.Coffee),
		nextID:  1,
	}
}

// GetAllCoffees retrieves all coffees from the in-memory storage
func (r *InMemoryCoffeeRepository) GetAllCoffees() ([]model.Coffee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	coffees := make([]model.Coffee, 0, len(r.coffees))
	for _, coffee := range r.coffees {
		coffees = append(coffees, coffee)
	}
	return coffees, nil
}

// GetCoffeeByID retrieves a coffee by its ID
func (r *InMemoryCoffeeRepository) GetCoffeeByID(id int64) (*model.Coffee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	coffee, exists := r.coffees[id]
	if !exists {
		return nil, errors.New("coffee not found")
	}
	return &coffee, nil
}

// CreateCoffee adds a new coffee to the in-memory storage
func (r *InMemoryCoffeeRepository) CreateCoffee(coffee *model.Coffee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	coffee.ID = r.nextID
	coffee.CreatedAt = time.Now().UTC()
	coffee.UpdatedAt = coffee.CreatedAt
	r.coffees[coffee.ID] = *coffee
	r.nextID++
	return nil
}

// UpdateCoffee updates an existing coffee in the in-memory storage
func (r *InMemoryCoffeeRepository) UpdateCoffee(coffee *model.Coffee) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.coffees[coffee.ID]; !exists {
		return errors.New("coffee not found")
	}

	coffee.UpdatedAt = time.Now().UTC()
	r.coffees[coffee.ID] = *coffee
	return nil
}

// DeleteCoffee removes a coffee from the in-memory storage
func (r *InMemoryCoffeeRepository) DeleteCoffee(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.coffees[id]; !exists {
		return errors.New("coffee not found")
	}

	delete(r.coffees, id)
	return nil
}
