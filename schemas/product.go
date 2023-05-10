package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       float64
	Description string
	Quantity    int
}

type ProductResponse struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
	Name        string     `json:"name"`
	Price       float64    `json:"price"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
}
