package main

import (
	"fmt"
	"errors"
)

type Matematica struct{}

func (m Matematica) Fatorial(n int) (int64, error) {
	if n < 0 {
		return 0, errors.New("o fatorial não está definido para números negativos")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}
	resultado := int64(1)
	for i := 2; i <= n; i++ {
		resultado *= int64(i)
	}
	return resultado, nil
}

func main() {
	mat := Matematica{}
	numero := 5

	fatorial, err := mat.Fatorial(numero)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O fatorial de %d é %d\n", numero, fatorial)
	}
}
