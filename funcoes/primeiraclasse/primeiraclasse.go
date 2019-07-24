package main

import "fmt"

// soma e sub são funções anonimas

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 2))
}

var soma = func(a, b int) int {
	return a + b
}
