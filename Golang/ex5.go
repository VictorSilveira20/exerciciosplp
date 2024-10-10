package main

import "fmt"

type Animal interface {
	Som() string
}

type Cachorro struct {
	nome string
}

func (c Cachorro) Som() string {
	return fmt.Sprintf("%s faz: Au Au!", c.nome)
}

type Gato struct {
	nome string
}

func (g Gato) Som() string {
	return fmt.Sprintf("%s faz: Miau!", g.nome)
}

func emitirSons(animais []Animal) {
	for _, animal := range animais {
		fmt.Println(animal.Som())
	}
}

func main() {
	animais := []Animal{
		Cachorro{"Valente"},
		Gato{"Marley"},
		Cachorro{"Sansão"},
		Gato{"Luna"},
	}

	emitirSons(animais)
}
