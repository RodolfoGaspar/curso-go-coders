package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// struct p/ json
	p1 := produto{1, "Notebook", 1899.78, []string{"promoção", "eletronico"}}

	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	// json p/ struct
	var p2 produto
	jsonString := `{"id":2, "nome": "caneta", "preco":8.9, "tags": ["papelaria", "importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}
