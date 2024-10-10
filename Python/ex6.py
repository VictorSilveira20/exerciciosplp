class Motor:
    def __init__(self, cilindradas, tipo_combustivel):
        self.cilindradas = cilindradas
        self.tipo_combustivel = tipo_combustivel

    def detalhes_motor(self):
        return f"Motor: {self.cilindradas}cc, Combustível: {self.tipo_combustivel}"


class Carro:
    def __init__(self, marca, modelo, ano, motor):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.motor = motor

    def detalhes(self):
        return f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}, {self.motor.detalhes_motor()}"


motor_teste = Motor(2000, "Gasolina")

carro_teste = Carro("Honda", "Civic", 2020, motor_teste)

print(carro_teste.detalhes())
