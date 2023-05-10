package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Street       string
	BuildingName string
	Locality     string
	City         string
	State        string
	PinCode      string
	// Customer     Customer
}

type AddressResponse struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
	Street       string     `json:"Street"`
	BuildingName string     `json:"BuildingName"`
	Locality     string     `json:"Locality"`
	City         string     `json:"City"`
	State        string     `json:"State"`
	PinCode      string     `json:"PinCode"`
	Customer     Customer   `json:"Customer"`
}
