package models

type Tutor struct {
	ID       int    `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Foto     string `json:"foto"`
	Nome     string `json:"nome"`
	Telefone string `json:"telefone"`
	Cidade   string `json:"cidade"`
	Sobre    string `json:"sobre"`
}
