package repository

import (
	"errors"
	"sync"

	"coffee-shop-api/internal/model"
)

var (
	ErrCoffeeNotFound = errors.New("coffee not found")
)

// CoffeeRepository defines the interface for coffee data operations
type CoffeeRepository interface {
	GetAll() []model.Coffee
	GetByID(id int) (model.Coffee, error)
	Create(coffee model.Coffee) model.Coffee
	Update(id int, coffee model.Coffee) (model.Coffee, error)
	Delete(id int) error
}

// InMemoryCoffeeRepository implements CoffeeRepository with an in-memory store
type InMemoryCoffeeRepository struct {
	coffees map[int]model.Coffee
	mu      sync.RWMutex
	nextID  int
}

// NewInMemoryCoffeeRepository creates a new instance of InMemoryCoffeeRepository
func NewInMemoryCoffeeRepository() *InMemoryCoffeeRepository {
	return &InMemoryCoffeeRepository{
		coffees: make(map[int]model.Coffee),
		nextID:  1,
	}
}

func (r *InMemoryCoffeeRepository) GetAll() []model.Coffee {
	r.mu.RLock()
	defer r.mu.RUnlock()

	coffees := make([]model.Coffee, 0, len(r.coffees))
	for _, coffee := range r.coffees {
		coffees = append(coffees, coffee)
	}
	return coffees
}

func (r *InMemoryCoffeeRepository) GetByID(id int) (model.Coffee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	coffee, exists := r.coffees[id]
	if !exists {
		return model.Coffee{}, ErrCoffeeNotFound
	}
	return coffee, nil
}

func (r *InMemoryCoffeeRepository) Create(coffee model.Coffee) model.Coffee {
	r.mu.Lock()
	defer r.mu.Unlock()

	coffee.ID = r.nextID
	r.nextID++
	r.coffees[coffee.ID] = coffee
	return coffee
}

func (r *InMemoryCoffeeRepository) Update(id int, coffee model.Coffee) (model.Coffee, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.coffees[id]; !exists {
		return model.Coffee{}, ErrCoffeeNotFound
	}

	coffee.ID = id
	r.coffees[id] = coffee
	return coffee, nil
}

func (r *InMemoryCoffeeRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.coffees[id]; !exists {
		return ErrCoffeeNotFound
	}

	delete(r.coffees, id)
	return nil
}
