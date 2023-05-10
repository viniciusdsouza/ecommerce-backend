package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderDate    time.Time
	OrderStatus  string
	Total        float64
	// CreditNumber CreditCard
	customer     Customer
	OrderedItems []CartItem
	OrderAddress Address
}

type OrderResponse struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	OrderDate    time.Time  `json:"orderDate"`
	OrderStatus  string     `json:"orderStatus"`
	Total        float64    `json:"total"`
	// CreditNumber CreditCard `json:"creditNumber"`
	Customer     Customer   `json:"customer"`
	OrderedItems []CartItem `json:"orderedItems"`
	OrderAddress Address    `json:"orderAddress"`
}
