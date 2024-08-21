package main

import "fmt"

var _ Greeter = alice{}
var _ Greeter = (*bob)(nil)

type Greeter interface {
	Hello()
}

type alice struct{}

func (alice) Hello() {
	fmt.Println("Eu sou Alice")
}

type bob struct{}

func (*bob) Hello() {
	fmt.Println("Eu sou Bob")
}

func main() {
	p := []interface{}{alice{}, bob{}, &alice{}, &bob{}}
	for _, person := range p {
		switch greeter := person.(type) {
		case Greeter:
			greeter.Hello()
		default:
			fmt.Println("Nao posso dizer ola")
		}
	}
}
