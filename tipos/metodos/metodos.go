package main

import (
	"fmt"
	"strings"
)

func main() {
	p1 := pessoa{"Pedro", "Rocha"}
	fmt.Println(p1.getNomeCompeto())

	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompeto())
}

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompeto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}
