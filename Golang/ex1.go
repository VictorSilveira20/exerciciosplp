package main

import (
	"fmt"
)

type Carro struct {
	marca  string
	modelo string
	ano    int
}

func (c Carro) detalhes() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d", c.marca, c.modelo, c.ano)
}

func main() {
	carro1 := Carro{"Toyota", "Corolla", 2020}
	carro2 := Carro{"Honda", "Civic", 2018}
	carro3 := Carro{"Ford", "Mustang", 2022}

	fmt.Println(carro1.detalhes())
	fmt.Println(carro2.detalhes())
	fmt.Println(carro3.detalhes())
}
