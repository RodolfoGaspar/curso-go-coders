package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Geraldo": 123.34,
			"Gustavo": 288.12,
		},
		"J": {
			"Joao": 292.39,
			"Jose": 197.28,
		},
		"P": {
			"Pedro": 890.37,
			"Paulo": 639.45,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
