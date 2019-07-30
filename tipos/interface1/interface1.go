package main

import "fmt"

func main() {
	var coisa imprimivel = pessoa{"Rodolfo", "Gaspar"}
	imprimir(coisa)

	coisa = produto{"roda", 1.23}
	imprimir(coisa)
}

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}
