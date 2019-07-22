package toy

import (
	"time"
)

// Toy model struct for encapsulation
type Toy struct {
	ID        string    `json:"id" db:"id"`
	Owner     string    `json:"owner" db:"owner"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
	DeletedAt time.Time `json:"-" db:"deletedAt"`
}
