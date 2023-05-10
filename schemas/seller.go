package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     float64
	Password  string
	Products  []Product
}

type SellerResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     float64    `json:"email"`
	Password  string     `json:"password"`
	Products  []Product  `json:"products"`
}
