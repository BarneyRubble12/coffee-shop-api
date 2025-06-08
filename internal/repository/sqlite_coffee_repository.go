package repository

import (
	"database/sql"
	"errors"
	"time"

	"coffee-shop-api/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteCoffeeRepository implements CoffeeRepository interface using SQLite
type SQLiteCoffeeRepository struct {
	db *sql.DB
}

// NewSQLiteCoffeeRepository creates a new SQLite repository instance
func NewSQLiteCoffeeRepository(dbPath string) (*SQLiteCoffeeRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create the coffees table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS coffees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteCoffeeRepository{db: db}, nil
}

// GetAllCoffees retrieves all coffees from the database
func (r *SQLiteCoffeeRepository) GetAllCoffees() ([]model.Coffee, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price, created_at, updated_at 
		FROM coffees
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []model.Coffee
	for rows.Next() {
		var coffee model.Coffee
		var createdAt, updatedAt string
		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Description,
			&coffee.Price,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		coffee.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		coffee.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
		coffees = append(coffees, coffee)
	}

	return coffees, nil
}

// GetCoffeeByID retrieves a coffee by its ID
func (r *SQLiteCoffeeRepository) GetCoffeeByID(id int64) (*model.Coffee, error) {
	var coffee model.Coffee
	var createdAt, updatedAt string

	err := r.db.QueryRow(`
		SELECT id, name, description, price, created_at, updated_at 
		FROM coffees 
		WHERE id = ?
	`, id).Scan(
		&coffee.ID,
		&coffee.Name,
		&coffee.Description,
		&coffee.Price,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("coffee not found")
		}
		return nil, err
	}

	coffee.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	coffee.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt)
	return &coffee, nil
}

// CreateCoffee adds a new coffee to the database
func (r *SQLiteCoffeeRepository) CreateCoffee(coffee *model.Coffee) error {
	now := time.Now().UTC()
	coffee.CreatedAt = now
	coffee.UpdatedAt = now

	result, err := r.db.Exec(`
		INSERT INTO coffees (name, description, price, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, coffee.Name, coffee.Description, coffee.Price, coffee.CreatedAt.Format(time.RFC3339), coffee.UpdatedAt.Format(time.RFC3339))
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	coffee.ID = id

	return nil
}

// UpdateCoffee updates an existing coffee in the database
func (r *SQLiteCoffeeRepository) UpdateCoffee(coffee *model.Coffee) error {
	coffee.UpdatedAt = time.Now().UTC()

	result, err := r.db.Exec(`
		UPDATE coffees 
		SET name = ?, description = ?, price = ?, updated_at = ?
		WHERE id = ?
	`, coffee.Name, coffee.Description, coffee.Price, coffee.UpdatedAt.Format(time.RFC3339), coffee.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("coffee not found")
	}

	return nil
}

// DeleteCoffee removes a coffee from the database
func (r *SQLiteCoffeeRepository) DeleteCoffee(id int64) error {
	result, err := r.db.Exec("DELETE FROM coffees WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("coffee not found")
	}

	return nil
}

// Close closes the database connection
func (r *SQLiteCoffeeRepository) Close() error {
	return r.db.Close()
}
