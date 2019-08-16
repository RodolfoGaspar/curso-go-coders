package main

import (
	"fmt"
	"time"
)

func main() {
	c := juntar(falar("Joao"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}

// internamente este metodo irá funcionar como um generator
func falar(pessoa string) <-chan string {
	// aqui será encapsulado a criação do canal, a chamada da go routine e será retornado um canal somente leitura
	c := make(chan string)
	// função anonima
	// será executada como um go routine
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i)
		}
	}()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}
