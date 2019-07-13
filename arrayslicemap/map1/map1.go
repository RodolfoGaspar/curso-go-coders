package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "maria"
	aprovados[12121212121] = "ana"
	aprovados[12121212321] = "joana"
	aprovados[12345578978] = "mariana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12121212121])

	delete(aprovados, 12121212121)
	fmt.Println(aprovados[12121212121])
}
