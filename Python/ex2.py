class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0

    def detalhes(self):
        return f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}"

    def acelerar(self, incremento):
        self.velocidade += incremento
        return f"{self.marca} {self.modelo} acelerou. Velocidade atual: {self.velocidade} km/h"

    def frear(self, decremento):
        self.velocidade = max(0, self.velocidade - decremento)
        return f"{self.marca} {self.modelo} freou. Velocidade atual: {self.velocidade} km/h"

    def exibir_velocidade(self):
        return f"Velocidade atual do {self.marca} {self.modelo}: {self.velocidade} km/h"


carro_teste = Carro("Toyota", "Corolla", 2020)

print(carro_teste.acelerar(50))
print(carro_teste.frear(20))
print(carro_teste.exibir_velocidade())