class Animal:
    def __init__(self, nome):
        self.nome = nome


class Cachorro(Animal):
    def som(self):
        return f"{self.nome} faz: Au Au!"


class Gato(Animal):
    def som(self):
        return f"{self.nome} faz: Miau!"


def emitir_sons(animais):
    for animal in animais:
        print(animal.som())


animais = [Cachorro("Valente"), Gato("Marley"), Cachorro("Sansão"), Gato("Luna")]

emitir_sons(animais)
