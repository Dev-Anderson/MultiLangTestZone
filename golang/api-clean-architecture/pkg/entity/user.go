package entity

import "time"

// defina a estrutura de dados basicas que representam  o nucleo do seu dominio
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	BirthDate time.Time `json:"birthdate"`
}
