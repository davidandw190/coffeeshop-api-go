package services

import (
	"context"
	"time"
)

type Coffee struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Roast     string    `json:"roast"`
	Image     string    `json:"image"`
	Region    string    `json:"region"`
	Price     float64   `json:"price"`
	GrindUnit int16     `json:"grind_unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllCoffees retrieves all coffee products from the database.
func (c *Coffee) GetAllCoffees() ([]*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	SELECT id, name, image, roast, region, price, grind_unit, created_at, updated_at
	FROM coffees
	`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []*Coffee
	for rows.Next() {
		var coffee Coffee
		if err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Image,
			&coffee.Roast,
			&coffee.Region,
			&coffee.Price,
			&coffee.GrindUnit,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		); err != nil {
			return nil, err
		}

		coffees = append(coffees, &coffee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return coffees, nil
}

// CreateCoffee inserts a new coffee product into the database.
func (c *Coffee) CreateCoffee(coffee Coffee) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        INSERT INTO coffees(name, image, region, roast, price, grind_unit, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `

	var id string

	err := db.QueryRowContext(
		ctx,
		query,
		coffee.Name,
		coffee.Image,
		coffee.Region,
		coffee.Roast,
		coffee.Price,
		coffee.GrindUnit,
		time.Now(),
		time.Now(),
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	coffee.ID = id

	return &coffee, nil
}

func (c *Coffee) GetCoffeeByID(id string) (*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT id, name, image, roast, region, price, grind_unit, created_at, updated_at 
        FROM coffees
        WHERE id = $1
    `
	var coffee Coffee

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&coffee.ID,
		&coffee.Name,
		&coffee.Image,
		&coffee.Roast,
		&coffee.Region,
		&coffee.Price,
		&coffee.GrindUnit,
		&coffee.CreatedAt,
		&coffee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &coffee, nil
}
