package main

import (
	"flag"
	"fmt"
)

var p1, p2, p3, p4 string

func init() {
	flag.StringVar(&p1, "parametro1", "valor_padrao_parametro1", "Primeiro parametro de configuracao")
	flag.StringVar(&p2, "parametro2", "valor_padrao_parametro2", "Segundo parametro de configuracao")
	flag.StringVar(&p3, "parametro3", "valor_padrao_parametro3", "Terceiro parametro de configuracao")
	flag.StringVar(&p4, "parametro4", "valor_padrao_parametro4", "Quarto parametro de configuracao")
}

func main() {
	flag.Parse()

	fmt.Printf("Parametro1: %s\n", p1)
	fmt.Printf("Parametro2: %s\n", p2)
	fmt.Printf("Parametro3: %s\n", p3)
	fmt.Printf("Parametro4: %s\n", p4)
}
