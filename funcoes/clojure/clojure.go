package main

import "fmt"

// em go, uma função tb é um clojure

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := clojure()
	imprimeX()
}

func clojure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}
