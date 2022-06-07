package entities

import "time"

type BaseEntity struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
