class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro):
        if isinstance(outro, Produto):
            return Produto("Soma de " + self.nome + " e " + outro.nome, self.preco + outro.preco)
        return NotImplemented

    def __str__(self):
        return f"{self.nome}: R${self.preco:.2f}"


def main():
    produto1 = Produto("Pastel", 50.0)
    produto2 = Produto("Coxinha", 30.0)

    produto_soma = produto1 + produto2
    print(produto_soma)

if __name__ == "__main__":
    main()
