package reposity

import (
	"database/sql"
	"duas-conexao-gorm-sq/schema"

	"gorm.io/gorm"
)

func CreatePedidoGorm(db *gorm.DB, novoPedido schema.Pedido) error {
	err := db.Create(&novoPedido)
	if err != nil {
		return err.Error
	}
	return nil
}

func GetPedidosClienteGorm(db *gorm.DB, clienteID int) ([]schema.Pedido, error) {
	var pedidos []schema.Pedido
	err := db.Where("cliente_id = ?", clienteID).Find(&pedidos)
	if err != nil {
		return nil, err.Error
	}
	return pedidos, nil
}

func GetPedidoSQL(db *sql.DB) ([]schema.Pedido, error) {
	var pedidos []schema.Pedido
	rows, err := db.Query(`
		select * from  pedidos
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var cliente_id int
		var produto_id int
		var quantidade int
		err = rows.Scan(&id, &cliente_id, &produto_id, &quantidade)
		if err != nil {
			return nil, err
		}
		pedidos = append(pedidos, schema.Pedido{ID: id, ClienteID: cliente_id, ProdutoID: produto_id, Quantidade: quantidade})
	}
	return pedidos, nil
}
