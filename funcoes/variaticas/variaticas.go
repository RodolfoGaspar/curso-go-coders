package main

import "fmt"

func main() {
	fmt.Printf("Média: %.2f\n", media(2, 3, 4))
	fmt.Printf("Média: %.2f\n", media(10, 2, 8, 99))
}

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}
