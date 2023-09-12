package models

type Produto struct {
	ID         uint    `json:"id"`
	Descricao  string  `json:"descricao"`
	Quantidade float64 `json:"quantidade"`
	Preco      float64 `json:"preco"`
}
