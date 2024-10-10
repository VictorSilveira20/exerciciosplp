class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def detalhes(self):
        return f"Nome: {self.nome}, Cargo: {self.cargo}, Salário: R${self.salario:.2f}"


class Empresa:
    def __init__(self, nome_empresa):
        self.nome_empresa = nome_empresa
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def listar_empregados(self):
        print(f"Empregados na empresa {self.nome_empresa}:")
        for empregado in self.empregados:
            print(empregado.detalhes())


empregado_1 = Empregado("Rennan", "Desenvolvedor", 8000)
empregado_2 = Empregado("Rodrigo", "Gerente de Projetos", 12000)
empregado_3 = Empregado("Guilherme", "Designer", 7000)

empresa = Empresa("DATAPREV")

empresa.adicionar_empregado(empregado_1)
empresa.adicionar_empregado(empregado_2)
empresa.adicionar_empregado(empregado_3)

empresa.listar_empregados()
