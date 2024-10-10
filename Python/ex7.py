class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)

    def listar_escolas(self):
        return [escola.nome for escola in self.escolas]


class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)
            professor.adicionar_escola(self)

    def listar_professores(self):
        return [professor.nome for professor in self.professores]


professor_1 = Professor("Dr. Ricardo")
professor_2 = Professor("Prof. Tatyana")

escola_1 = Escola("Escola Anglo")
escola_2 = Escola("Escola Marista")

escola_1.adicionar_professor(professor_1)
escola_1.adicionar_professor(professor_2)
escola_2.adicionar_professor(professor_2)

print(f"Professores na {escola_1.nome}: {escola_1.listar_professores()}")
print(f"Professores na {escola_2.nome}: {escola_2.listar_professores()}")

print(f"Escolas onde {professor_1.nome} leciona: {professor_1.listar_escolas()}")
print(f"Escolas onde {professor_2.nome} leciona: {professor_2.listar_escolas()}")
