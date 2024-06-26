package main

import (
	"fmt"
	"go/version"
	"slices"
)

func main() {
	old := "go1.21"
	t := "go1.x"
	new := "go1.22"

	for _, v := range []string{old, new, t} {
		if version.IsValid(v) {
			fmt.Printf("%s is valid\n", v)
		} else {
			fmt.Printf("%s is invalid]\n", v)
		}
	}

	switch version.Compare(old, new) {
	case 0:
		fmt.Printf("%s and %s are equal\n", old, new)
	case 1:
		fmt.Printf("%s is greater than %s\n", old, new)
	case -1:
		fmt.Printf("%s is less than %s\n", old, new)
	}

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	//add valores dentro de uma outra variavel
	join := slices.Concat(a, b)
	fmt.Println(join)
}
