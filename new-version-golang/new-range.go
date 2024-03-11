package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

//N retorna um numero pseuodoaleatorio no intervalo semiaberto[0,n] da fonte padrao.
//O parametro de tipo int pode ser qualquer tipo inteiro. Ele entra em panico se n <= 0.

func main() {
	for range 10 {
		d := rand.N(time.Second * 2)
		fmt.Println(d)
	}
}
