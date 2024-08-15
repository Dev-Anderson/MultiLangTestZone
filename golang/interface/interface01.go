package main

import "fmt"

// Define a interface Animal
type Animal interface {
	Falar() string
}

// Struct Cachorro que implementa a interface Animal
type Cachorro struct{}

// Método Falar para o tipo Cachorro
func (c Cachorro) Falar() string {
	return "Au Au!"
}

// Struct Gato que implementa a interface Animal
type Gato struct{}

// Método Falar para o tipo Gato
func (g Gato) Falar() string {
	return "Miau!"
}

func main() {
	var animal Animal

	// Cachorro satisfaz a interface Animal
	animal = Cachorro{}
	fmt.Println(animal.Falar()) // Saída: Au Au!

	// Gato satisfaz a interface Animal
	animal = Gato{}
	fmt.Println(animal.Falar()) // Saída: Miau!
}
