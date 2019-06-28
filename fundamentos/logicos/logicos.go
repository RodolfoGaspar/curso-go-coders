package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	tomarSorvete := trab1 || trab2

	return comprarTv32, comprarTv50, tomarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(false, true)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(true, false)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)

	tv50, tv32, sorvete = compras(false, false)
	fmt.Printf("\nTv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t",
		tv50, tv32, sorvete, !sorvete)
}
