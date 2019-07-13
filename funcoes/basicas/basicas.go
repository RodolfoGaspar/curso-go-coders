package main

import "fmt"

func main() {
	f1()
	f2("Hello", "world")
	fmt.Println(f3())
	fmt.Println(f4("Olá", "Mundo"))
	fmt.Println(f5())

	r3, r4 := f3(), f4("Olá", "Mundo")
	fmt.Println(r3, r4)

	r51, r52, r53 := f5()
	fmt.Println(r51, r52, r53)

	r51b, r52b, _ := f5()
	fmt.Println(r51b, r52b)

}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string, string) {
	return "F5:", "Go", "Lang"
}
