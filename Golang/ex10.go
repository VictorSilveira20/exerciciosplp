package main

import "fmt"

func somarDois(a int, b int) int {
	return a + b
}

func somarTres(a int, b int, c int) int {
	return a + b + c
}

func main() {
	fmt.Println("Soma de 5 e 10:", somarDois(5, 10))
	fmt.Println("Soma de 5, 10 e 15:", somarTres(5, 10, 15))
}
