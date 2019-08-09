package main

import (
	"fmt"
	"time"
)

// chanel (canal) é a forma de comunicação entre as goroutines
// chanel é um tipo, assim como int, float, etc

func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 4 * base //enviando dados para o canal
}
