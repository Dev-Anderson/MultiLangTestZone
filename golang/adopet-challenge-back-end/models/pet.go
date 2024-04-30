package models

type Pet struct {
	ID             int    `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Nome           string `json:"nome"`
	Idade          uint   `json:"idade"`
	Porte          string `json:"porte"`
	Caracteristica string `json:"caracteristica"`
	Cidade         string `json:"cidade"`
	UF             string `json:"uf"`
	AbrigoID       int    `json:"abrigoid" swaggerignore:"true"`
	Abrigo         Abrigo `gorm:"references:ID"`
}
