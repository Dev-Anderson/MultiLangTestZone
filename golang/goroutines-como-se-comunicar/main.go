package main

import (
	"fmt"
	"time"
)

func minhaGoroutine(canal chan string) {
	time.Sleep(time.Second)
	canal <- "Ola, goroutine!"
}

func main() {
	meuCanal := make(chan string) //indicando o que sera enviado para o canal

	go minhaGoroutine(meuCanal) //inicia a goroutine

	mensagem := <-meuCanal //recebe a mensagem do canal
	fmt.Println(mensagem)
}
