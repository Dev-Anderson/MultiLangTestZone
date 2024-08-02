package main

import "fmt"

func asdf(a int) {
	fmt.Println(a)
}

func main() {
	a := 10
	defer asdf(a)
	a = 20
}

// output vai ser 10
// Isso porque quando voce usa a instrucao defer, ela captura os valores imediatamente. Isso e chamado de captura
// por valor de "a" e enviado para funca asdf() e definido como 10 quando o defer e agendado, mesmo que "a" haja
// alteracoes posteriores.
