class SaldoInsuficienteException(Exception):
    def __init__(self, mensagem):
        super().__init__(mensagem)


class ContaBancaria:
    def __init__(self, titular, saldo_inicial=0):
        self.titular = titular
        self.saldo = saldo_inicial

    def depositar(self, valor):
        if valor <= 0:
            raise ValueError("O valor do depósito deve ser positivo.")
        self.saldo += valor

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException("Saldo insuficiente para realizar o saque.")
        self.saldo -= valor

    def obter_saldo(self):
        return self.saldo


def main():
    conta = ContaBancaria("Coisinha", 100)

    try:
        conta.sacar(150)
    except SaldoInsuficienteException as e:
        print(e)


if __name__ == "__main__":
    main()
