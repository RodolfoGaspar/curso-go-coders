package main

import (
	"fmt"
	"time"

	"github.com/rodolfogaspar/html"
)

func main() {
	campeao := oMaisRapido(
		"https://www.amazon.com",
		"https://devsontable.com",
		"https://www.youtube.com",
	)
	fmt.Println(campeao)
}

func oMaisRapido(url1, url2, url3 string) string {
	// c1 := hmtl.Titulo(url1)
	// c0 := html.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "todos perderam"
		// default:
		// 	return "sem resposta"
	}
}
