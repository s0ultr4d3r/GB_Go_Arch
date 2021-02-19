package models

import "time"

type Item struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"crated_at"`
	UpdatedAt time.Time `json:'updated_at"`
}
