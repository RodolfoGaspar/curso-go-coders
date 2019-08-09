package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //canal sem buffer
	fmt.Println("make")
	go rotina(c)
	fmt.Println("go")
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c)
	fmt.Println("Fim")
}

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido")
}
