package entity

import (
	"time"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	NameBank  string
	OwnerName string
}

type BankResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	NameBank  string    `json:"nameBank"`
	OwnerName string    `json:"ownerName"`
}
