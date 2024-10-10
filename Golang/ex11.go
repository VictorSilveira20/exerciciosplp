package main

import "fmt"

type Funcionario interface {
	CalcularSalario() float64
}

type FuncionarioHorista struct {
	nome            string
	valorHora       float64
	horasTrabalhadas int
}

func (f FuncionarioHorista) CalcularSalario() float64 {
	return f.valorHora * float64(f.horasTrabalhadas)
}

type FuncionarioAssalariado struct {
	nome          string
	salarioMensal float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
	return f.salarioMensal
}

func main() {
	horista := FuncionarioHorista{"Sarah", 50.0, 160}
	assalariado := FuncionarioAssalariado{"Rodrigo", 3000.0}

	fmt.Printf("%s - Salário: R$%.2f\n", horista.nome, horista.CalcularSalario())
	fmt.Printf("%s - Salário: R$%.2f\n", assalariado.nome, assalariado.CalcularSalario())
}
