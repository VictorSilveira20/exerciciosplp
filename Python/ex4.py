class Animal:
    def __init__(self, nome):
        self.nome = nome


class Cachorro(Animal):
    def som(self):
        return f"{self.nome} faz: Au Au!"


class Gato(Animal):
    def som(self):
        return f"{self.nome} faz: Miau!"


cachorro = Cachorro("Thor")
gato = Gato("Mel")

print(cachorro.som())
print(gato.som())
