class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.__saldo = saldo_inicial  # Atributo saldo encapsulado (privado)
        self.titular = titular

    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            return f"Depósito de R${valor:.2f} realizado. Saldo atual: R${self.__saldo:.2f}"
        else:
            return "Valor de depósito inválido!"

    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            return f"Saque de R${valor:.2f} realizado. Saldo atual: R${self.__saldo:.2f}"
        else:
            return "Saldo insuficiente ou valor de saque inválido!"

    def exibir_saldo(self):
        return f"Saldo atual de {self.titular}: R${self.__saldo:.2f}"


conta_teste = ContaBancaria("João Silva", 1000)

print(conta_teste.depositar(500))
print(conta_teste.sacar(300))
print(conta_teste.exibir_saldo())
