package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

//N retorna um nuemro pseuodoaleatorio no intervalo semiabert[0, n] da fonte padrao
// O parametro de tipo int pode ser qualquer tipo inteiro. Ele entrar em panico se n <= 0

func main() {
	for range 10 {
		d := rand.N(time.Second * 2)
		fmt.Println(d)
	}
}
