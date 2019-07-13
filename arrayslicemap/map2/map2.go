package main

import "fmt"

func main() {
	funcsESalario := map[string]float64{
		"Jose":  500.00,
		"Joao":  602.34,
		"Maria": 1093.49,
	}

	fmt.Println(funcsESalario)

	funcsESalario["Ana"] = 789.37
	fmt.Println(funcsESalario)

	delete(funcsESalario, "Inexiste") //não acontece nada, nem erro

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}

}
