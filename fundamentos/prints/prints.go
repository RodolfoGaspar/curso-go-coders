package main

import (
	f "fmt"
)

func main() {
	f.Print("Mesma")
	f.Print(" Linha.")

	f.Println(" Nova")
	f.Println("Linha.")

	x := 3.141516

	// f.Println("O valor de x é " + x)
	xs := f.Sprint(x)
	f.Println("O valor de x é " + xs)
	f.Println("O valor de x é", x)

	f.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Ops"

	f.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	f.Printf("\n%v %v %v %v", a, b, c, d)
}
