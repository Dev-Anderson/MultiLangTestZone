package schema

type Pedido struct {
	ID         int
	ClienteID  int
	ProdutoID  int
	Quantidade int
}

type TotalVendas struct {
	Mes   int
	Total float64
}
