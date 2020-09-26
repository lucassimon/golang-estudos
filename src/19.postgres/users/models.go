package users

import (
	"time"
)

// User User schema of the user table
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int16     `json:"age"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
