package main

import "fmt"

func main() {
	var b esportivoLuxuoso = bmwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// pode adicionar outros metodos
}

type bmwm7 struct{}

func (b bmwm7) ligarTurbo() {
	fmt.Println("Turbo Ligado!")
}

func (b bmwm7) fazerBaliza() {
	fmt.Println("Fazendo Baliza!")
}
