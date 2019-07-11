package main

import "fmt"

func ImprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota ", nota)
	} else {
		fmt.Println("Reprovado com a nota ", nota)
	}
}

func main() {
	ImprimirResultado(7.3)
	ImprimirResultado(4.2)
}
