package main

import "fmt"

// map, reduce e filter são bons exemplos de utilização de função como parametro

func main() {

	fmt.Println(exec(multiplicacao, 2, 3))
}

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
	