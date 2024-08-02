package main

import "fmt"

func b() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	a()
}

func a() {
	defer fmt.Println(3)
	defer fmt.Println(4)
}

func main() {
	b()
}
