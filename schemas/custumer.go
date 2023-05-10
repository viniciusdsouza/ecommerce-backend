package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        float64
	Password     string
	Orders       []Order
	CustomerCart Cart
	// CardNumber   CreditCard
	Addresses Address
}

type CustomerResponse struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	FirstName    string     `json:"FirstName"`
	LastName     string     `json:"LastName"`
	Email        float64    `json:"Email"`
	Password     string     `json:"Password"`
	Orders       []Order    `json:"Orders"`
	CustomerCart Cart       `json:"CustomerCart"`
	// CardNumber   CreditCard `json:"CardNumber"`
	Addresses Address `json:"Addresses"`
}
