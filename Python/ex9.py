from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def imprimir(self):
        print("Imprimindo o Relatório.")

class Contrato(Imprimivel):
    def imprimir(self):
        print("Imprimindo o Contrato.")

def main():
    relatorio = Relatorio()
    contrato = Contrato()

    relatorio.imprimir()
    contrato.imprimir()

if __name__ == "__main__":
    main()
