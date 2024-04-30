package models

type Abrigo struct {
	ID       int    `json:"id" gorm:"primaryKey" swaggerignore:"true"`
	Foto     string `json:"foto"`
	Telefone string `json:"telefone"`
	Nome     string `json:"nome"`
	Cidade   string `json:"cidade"`
	UF       string `json:"uf"`
	Sobre    string `json:"sobre"`
}
