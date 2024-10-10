package main

import (
	"errors"
	"fmt"
)

type SaldoInsuficienteError struct {
	Mensagem string
}

func (e *SaldoInsuficienteError) Error() string {
	return e.Mensagem
}

type ContaBancaria struct {
	titular string
	saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) error {
	if valor <= 0 {
		return errors.New("o valor do depósito deve ser positivo")
	}
	c.saldo += valor
	return nil
}

func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.saldo {
		return &SaldoInsuficienteError{Mensagem: "Saldo insuficiente para realizar o saque."}
	}
	c.saldo -= valor
	return nil
}

func (c *ContaBancaria) ObterSaldo() float64 {
	return c.saldo
}

func main() {
	conta := ContaBancaria{titular: "Coisinha Três", saldo: 100}

	err := conta.Sacar(150)
	if err != nil {
		fmt.Println(err)
	}
}
