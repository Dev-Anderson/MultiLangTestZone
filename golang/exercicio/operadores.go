package main

import "fmt"

//Tipos de operadores
// &&(e) = E verdade se ambas as afirmacoes forem verdadeiras
// ||(ou) = E verdade se pelo menos uma afirmacao for verdadeira
// !(nao) = So e verdade se a afirmacao for false

func CanFastAttack(kningthIsAwake bool) bool {
	//se cavaleiro esta acordado retorna true
	if kningthIsAwake != true {
		return true
	}
	return false
}

func main() {
	var k = true
	fmt.Println(CanFastAttack(k))
}
