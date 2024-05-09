package models

import "gorm.io/gorm"

type UF struct {
	gorm.Model
	Nome    string   `json:"nome"`
	Sigla   string   `json:"sigla"`
	Cidades []Cidade `gorm:"foreignKey:UFID"`
}
