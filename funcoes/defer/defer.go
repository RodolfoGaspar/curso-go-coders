package main

import "fmt"

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(600))
}

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")

	if numero > 5000 {
		fmt.Println("Valor Alto")
		return 5000
	}
	fmt.Println("Valor Baixo")
	return numero
}
