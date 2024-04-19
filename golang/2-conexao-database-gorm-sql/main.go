package main

import (
	"duas-conexao-gorm-sq/database"
	gormdb "duas-conexao-gorm-sq/database/gorm"
	"duas-conexao-gorm-sq/reposity"
	"duas-conexao-gorm-sq/schema"
	"fmt"
	"log"
)

func main() {
	gormDB, err := gormdb.Connect()
	if err != nil {
		log.Fatal(err)
	}

	sqlDb, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close(sqlDb)

	//insert novo pedido Gorm
	novoPedido := schema.Pedido{ClienteID: 123, ProdutoID: 1234, Quantidade: 10}
	err = reposity.CreatePedidoGorm(gormDB, novoPedido)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Novo pedido inserido com sucesso:", novoPedido)

	//consulta pedido por ID 123 com o gorm
	p, err := reposity.GetPedidosClienteGorm(gormDB, 123)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pedidos do cliente 123: ", p)

	//consulta utilizando o SQL puro
	pedidos, err := reposity.GetPedidoSQL(sqlDb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resultado dos pedidos utilizando o SQL PURO: ", pedidos)

}
