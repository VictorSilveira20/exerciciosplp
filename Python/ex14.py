class Configuracao:
    _instancia = None

    def __new__(cls):
        if cls._instancia is None:
            cls._instancia = super(Configuracao, cls).__new__(cls)
            cls._instancia.configuracoes = {}
        return cls._instancia

    def set_configuracao(self, chave, valor):
        self.configuracoes[chave] = valor

    def get_configuracao(self, chave):
        return self.configuracoes.get(chave, None)


def main():
    config1 = Configuracao()
    config1.set_configuracao("host", "localhost")

    config2 = Configuracao()
    config2.set_configuracao("port", 8080)

    print(config1.get_configuracao("host"))
    print(config2.get_configuracao("port"))
    print(config1 is config2)

if __name__ == "__main__":
    main()
