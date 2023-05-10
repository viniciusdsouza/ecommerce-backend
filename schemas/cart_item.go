package schemas

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	Product  Product
	Quantity int
}

type CartItemResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Product   string     `json:"product"`
	Quantity  int        `json:"quantity"`
}
