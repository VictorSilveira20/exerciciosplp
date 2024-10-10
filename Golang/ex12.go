package main

import "fmt"

type Produto struct {
	nome  string
	preco float64
}

func (p Produto) Somar(outro Produto) Produto {
	return Produto{
		nome:  "Soma de " + p.nome + " e " + outro.nome,
		preco: p.preco + outro.preco,
	}
}

func main() {
	produto1 := Produto{"Nescau", 50.0}
	produto2 := Produto{"Toddy", 30.0}

	produtoSoma := produto1.Somar(produto2)
	fmt.Printf("%s: R$%.2f\n", produtoSoma.nome, produtoSoma.preco)
}
