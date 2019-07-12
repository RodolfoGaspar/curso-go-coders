package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	notaParaConceito(9.5)
	notaParaConceito(8.2)
	notaParaConceito(6.3)
	notaParaConceito(4.7)
	notaParaConceito(0.9)

}

func notaParaConceito(n float64) {
	switch {
	case n >= 9 && n <= 10:
		fmt.Println("A")
	case n >= 8 && n < 9:
		fmt.Println("B")
	case n >= 5 && n < 8:
		fmt.Println("C")
	case n >= 3 && n < 5:
		fmt.Println("D")
	default:
		fmt.Println("E")

	}
}
