package main

import (
	"fmt"
)

type Carro struct {
	marca     string
	modelo    string
	ano       int
	velocidade int
}

func (c Carro) detalhes() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d", c.marca, c.modelo, c.ano)
}

func (c *Carro) acelerar(incremento int) string {
	c.velocidade += incremento
	return fmt.Sprintf("%s %s acelerou. Velocidade atual: %d km/h", c.marca, c.modelo, c.velocidade)
}

func (c *Carro) frear(decremento int) string {
	if decremento > c.velocidade {
		c.velocidade = 0
	} else {
		c.velocidade -= decremento
	}
	return fmt.Sprintf("%s %s freou. Velocidade atual: %d km/h", c.marca, c.modelo, c.velocidade)
}

func (c Carro) exibirVelocidade() string {
	return fmt.Sprintf("Velocidade atual do %s %s: %d km/h", c.marca, c.modelo, c.velocidade)
}

func main() {
	carroTeste := Carro{"Toyota", "Corolla", 2020}

	fmt.Println(carroTeste.acelerar(50))
	fmt.Println(carroTeste.frear(20))
	fmt.Println(carroTeste.exibirVelocidade())
}
