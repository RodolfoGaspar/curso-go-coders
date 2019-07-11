package main

import "fmt"

func main() {
	i := 1 //ocupa 64bit ou 8bytes na memoria

	var p *int = nil

	p = &i // atribuindo o endereço na memoria da variavel i ao ponteiro p

	// Go não aritmetica de ponteiro
	// p++

	fmt.Println("i: ", i)
	fmt.Println("&i: ", &i)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)

	*p++ // desreferenciando (obtendo o valor)

	fmt.Println("i: ", i)
	fmt.Println("&i: ", &i)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)

	i++

	fmt.Println("i: ", i)
	fmt.Println("&i: ", &i)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
}
