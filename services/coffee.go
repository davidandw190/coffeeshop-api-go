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

func (c *Coffee) GetAllCoffees() ([]*Coffee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

}