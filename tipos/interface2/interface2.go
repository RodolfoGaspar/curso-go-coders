package main

import "fmt"

func main() {
	carro1 := ferrari{"F40", false, 0}
	fmt.Println(carro1)
	carro1.ligarTurbo()
	fmt.Println(carro1)

	var carro2 esportivo = &ferrari{"360 Modena", false, 0}
	fmt.Println(carro1, carro2)
	carro2.ligarTurbo()
	fmt.Println(carro1, carro2)
}

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}
