package main

import "fmt"

func main() {
	// slice com tamanho e capacidade 10
	s := make([]int, 10)
	s[9] = 12
	// fmt.Println(s)
	myPrint(s)

	// slice com tamanho 10 e capacidade 20
	s = make([]int, 10, 20)
	s[7] = 13
	// fmt.Println(s, len(s), cap(s))
	myPrint(s)

	// slice com tamanho e capacidade 20
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	// fmt.Println(s, len(s), cap(s))
	myPrint(s)

	// slice com tamanho 20 e capacidade 40
	s = append(s, 1)
	// fmt.Println(s, len(s), cap(s))
	myPrint(s)
}

func myPrint(s []int) {
	fmt.Println("slice: ", s, "len:", len(s), "cap: ", cap(s))
	// fmt.Println("ponteiro:", &s)
	for i, num := range s {
		fmt.Println(i, ":", &num, " ")
	}
}
