package main

import "fmt"

type ContaBancaria struct {
	saldo   float64
	titular string
}

func NewContaBancaria(titular string, saldoInicial float64) *ContaBancaria {
	return &ContaBancaria{saldo: saldoInicial, titular: titular}
}

func (c *ContaBancaria) Depositar(valor float64) string {
	if valor > 0 {
		c.saldo += valor
		return fmt.Sprintf("Depósito de R$%.2f realizado. Saldo atual: R$%.2f", valor, c.saldo)
	}
	return "Valor de depósito inválido!"
}

func (c *ContaBancaria) Sacar(valor float64) string {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		return fmt.Sprintf("Saque de R$%.2f realizado. Saldo atual: R$%.2f", valor, c.saldo)
	}
	return "Saldo insuficiente ou valor de saque inválido!"
}

func (c *ContaBancaria) ExibirSaldo() string {
	return fmt.Sprintf("Saldo atual de %s: R$%.2f", c.titular, c.saldo)
}

func main() {
	contaTeste := NewContaBancaria("João Silva", 1000)
	fmt.Println(contaTeste.Depositar(500))
	fmt.Println(contaTeste.Sacar(300))
	fmt.Println(contaTeste.ExibirSaldo())
}
