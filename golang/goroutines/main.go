package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()              //gerando uma variavel com a data e hora atual
	var wg sync.WaitGroup            //criando um grupo para as goroutines
	wg.Add(3)                        //passando o numero de goroutines
	retornos := make(chan string, 3) // criando um canal do tipo string que vai receber 3 valores

	go buscarUsuario(&wg, retornos)
	go buscarCartao(&wg, retornos)
	go buscarCompra(&wg, retornos)
	wg.Wait()       //esperando as goroutines rodar
	close(retornos) //fechando os canais

	//percorrendo os valores carregados dentro do canal
	for valor := range retornos {
		fmt.Println(valor)
	}

	fmt.Println("Demorou: ", time.Since(start))

}

func buscarUsuario(wg *sync.WaitGroup, retorno chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	retorno <- "Anderson"
}

func buscarCartao(wg *sync.WaitGroup, retorno chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 150)
	retorno <- "Mastercard"
}

func buscarCompra(wg *sync.WaitGroup, retorno chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	retorno <- "Compra..."
}
