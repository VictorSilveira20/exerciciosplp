package main

import (
	"fmt"
)

// Interface Imprimivel
type Imprimivel interface {
	Imprimir()
}

type Relatorio struct{}

func (r Relatorio) Imprimir() {
	fmt.Println("Imprimindo o Relatório.")
}

type Contrato struct{}

func (c Contrato) Imprimir() {
	fmt.Println("Imprimindo o Contrato.")
}

func main() {
	var relatorio Imprimivel = Relatorio{}
	var contrato Imprimivel = Contrato{}

	relatorio.Imprimir()
	contrato.Imprimir()
}
