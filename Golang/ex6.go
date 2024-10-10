package main

import "fmt"

type Motor struct {
	cilindradas     int
	tipoCombustivel string
}

func (m Motor) DetalhesMotor() string {
	return fmt.Sprintf("%dcc, Combustível: %s", m.cilindradas, m.tipoCombustivel)
}

type Carro struct {
	marca  string
	modelo string
	ano    int
	motor  Motor
}

func (c Carro) Detalhes() string {
	return fmt.Sprintf("Marca: %s, Modelo: %s, Ano: %d, Motor: %s", c.marca, c.modelo, c.ano, c.motor.DetalhesMotor())
}

func main() {
	motorTeste := Motor{2000, "Gasolina"}
	carroTeste := Carro{"Honda", "Civic", 2020, motorTeste}
	fmt.Println(carroTeste.Detalhes())
}
