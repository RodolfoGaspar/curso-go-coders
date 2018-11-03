package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	/// Cuidado...
	fmt.Println("teste " + string(123)) // o número entre parenteses, depois da string, representa o cód de um caracter na tab ASCII

	// int p string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)
}
