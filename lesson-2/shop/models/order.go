package models

import "time"

type Order struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	ItemIDs   []int32   `json:"item_ids"`
	CreatedAt time.Time `json:"crated_at"`
	UpdatedAt time.Time `json:'updated_at"`
}
