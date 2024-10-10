package main

import (
	"fmt"
	"sync"
)

type Configuracao struct {
	configuracao string
}

var (
	instancia *Configuracao
	once      sync.Once
)

func GetConfiguracao() *Configuracao {
	once.Do(func() {
		instancia = &Configuracao{configuracao: "Configuração padrão"}
	})
	return instancia
}

func (c *Configuracao) SetConfiguracao(configuracao string) {
	c.configuracao = configuracao
}

func (c *Configuracao) GetConfiguracao() string {
	return c.configuracao
}

func main() {
	config1 := GetConfiguracao()
	config1.SetConfiguracao("Host: localhost")

	config2 := GetConfiguracao()
	config2.SetConfiguracao("Port: 8080")

	fmt.Println(config1.GetConfiguracao())
	fmt.Println(config2.GetConfiguracao())
	fmt.Println(config1 == config2)
}
