package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(9.0))
	fmt.Println(notaParaConceito(8.3))
	fmt.Println(notaParaConceito(6.7))
	fmt.Println(notaParaConceito(4.6))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(12.1))
}

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough //continua descendo para os proximos cases
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "nota invalida"
	}
}
