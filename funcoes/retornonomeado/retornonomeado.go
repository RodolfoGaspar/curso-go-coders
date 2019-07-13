package main

import "fmt"

func main() {
	fmt.Println(trocar(1, 2))
	x, y := trocar(3, 4)
	fmt.Println(x, y)
}

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo
}
