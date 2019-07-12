package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if i := numeroAleatorio(); i > 5 { // tb é suportado no switch
		fmt.Println("Ganhou -", i)
	} else {
		fmt.Println("Perdeu -", i)
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
