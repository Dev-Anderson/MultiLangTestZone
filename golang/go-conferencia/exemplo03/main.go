package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch) //adicionado em outra thread
	consume(ch)
}

//aqui eu estou publicando uma mensagem
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i //publicando as coisas dentro do canal
	}
	close(ch) //fechando o canal
}

//pegando o que foi publicando
func consume(ch chan int) {

	for x := range ch {
		fmt.Println("Messagem processada:", x) //lendo a mensagem e depois processando ela....
	}
}

// neste exemplo ele publicou os numero de 1 a 10 utilizando a funcao "publish" e depois foi criado uma outra funcao
// que faz a leitura do canel e pega os dados que foram publicados
// dentro do "main" foi adicionado uma goroutine dentro da funcao "publish"  para conseguir fazer o processamento em
// outra thread e depois utilizado a funcao "consume" para captura o que esta dentro do canal
// obsrevacao veja que foi sempre passao do canal de comunicacao
