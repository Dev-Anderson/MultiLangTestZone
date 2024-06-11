package entity

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	NameCard   string
	NameOwner  string
	NumberCard int
	CvvCard    int
}

type CardResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	NameCard   string    `json:"nameCard"`
	NameOwner  string    `json:"nameOwner"`
	NumberCard string    `json:"numberCard"`
	CvvCard    string    `json:"cvvCard"`
}
