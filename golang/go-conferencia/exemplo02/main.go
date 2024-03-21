package main

import "fmt"

//Thread 01
func main() {
	canal := make(chan string)

	// Thread 02
	//funcao anonima
	go func() {
		canal <- "Golang Conference! - Vindo da Thread 02"
	}()

	msg := <-canal //se o canal tiver cheio. Esvazia ele!!!

	fmt.Println(msg) //pritando o que esta dentro o canal
}
