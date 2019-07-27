package main

import "fmt"

func main() {
	p := pedido{
		userId: 1,
		itens: []item{
			item{
				produtoId: 1,
				qtde:      2,
				preco:     12.10},
			item{
				produtoId: 2,
				qtde:      20,
				preco:     23.65},
			item{
				produtoId: 3,
				qtde:      15,
				preco:     14.4},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f\n", p.valorTotal())
}

type item struct {
	produtoId int
	qtde      int
	preco     float64
}

type pedido struct {
	userId int
	itens  []item
}

func (p pedido) valorTotal() float64 {

	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}
