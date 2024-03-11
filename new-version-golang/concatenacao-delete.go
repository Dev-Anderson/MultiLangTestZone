package main

import (
	"fmt"
	"slices"
)

func main() {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	//codigo para adicionar as duas variaveis dentro de um slice
	join := slices.Concat(a, b)
	//codigo para remover valores de um slice, repassando a posicao que deve ser removido neste caso posicao 0 ate a posicao 2
	join = slices.Delete(join, 0, 2)

	fmt.Println(join)

}
