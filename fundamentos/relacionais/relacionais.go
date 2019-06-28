package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings: ", "Banana" == "Banana") //igual
	fmt.Println("!=", 3 != 2)                      //diferente
	fmt.Println("<", 3 < 2)                        //menor que
	fmt.Println(">", 3 > 2)                        //maior que
	fmt.Println("<=", 3 <= 2)                      //menor igual a
	fmt.Println(">=", 3 >= 2)                      //maior igual a

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2) //Não compara referencias de memoria, mas sim o com o valor
	fmt.Println("Datas: ", d1.Equal(d2))
	fmt.Println(d1)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"joão"}

	fmt.Println("Pessoas:", p1 == p2)

}
