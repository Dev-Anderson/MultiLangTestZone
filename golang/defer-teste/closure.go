package main

import "fmt"

func asdf(a int) {
	fmt.Println(a)
}

func main() {
	a := 10
	defer func() {
		asdf(a)
	}()
	a = 20
}

// Funções Go podem ser closures. Um closure é um valor de função que referencia variávei de fora do seu corpo.
// A função pode acessar e atribuir às variáveis referenciadas; nesse sentido,  a função é vinculada às variaveis.
