package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	NameTransaction string
	Value           int
	Category        string
}

type TransactionResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"`
	NameTransaction string    `json:"nameTrasaction"`
	Value           int       `json:"value"`
	Category        string    `json:"category"`
}
