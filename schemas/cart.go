package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	CartItems []CartItem
	CartTotal float64
	// Customer  Customer
}

type CartResponse struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	CartItems []CartItem `json:"cartItems"`
	CartTotal float64    `json:"cartTotal"`
	Customer  Customer   `json:"customer"`
}
