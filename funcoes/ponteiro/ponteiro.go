package main

import "fmt"

func main() {
	n := 1
	fmt.Println(n)

	inc1(n) // por valor
	fmt.Println(n)

	// & => é usado p obter o enderenço da variável
	inc2(&n) // por referencia
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *

func inc2(n *int) {
	// revisão: = * é usado para desreferenciar
	// 				Ou seja, é usado p ter acesso ao valor no qual o ponteiro aponta

	*n++
}
