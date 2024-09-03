package entities

import "github.com/google/uuid"

// Address struct
type Address struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Line1   string    `json:"line1"`
	Line2   string    `json:"line2"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	ZipCode string    `json:"zip_code"`
	Country string    `json:"country"`
}
