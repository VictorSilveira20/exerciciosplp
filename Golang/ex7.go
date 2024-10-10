package main

import "fmt"

type Professor struct {
	nome    string
	escolas []Escola
}

func NewProfessor(nome string) *Professor {
	return &Professor{nome: nome, escolas: []Escola{}}
}

func (p *Professor) AdicionarEscola(escola *Escola) {
	for _, e := range p.escolas {
		if e.nome == escola.nome {
			return
		}
	}
	p.escolas = append(p.escolas, *escola)
	escola.AdicionarProfessor(p)
}

func (p *Professor) ListarEscolas() []string {
	nomes := []string{}
	for _, escola := range p.escolas {
		nomes = append(nomes, escola.nome)
	}
	return nomes
}

type Escola struct {
	nome      string
	professores []Professor
}

func NewEscola(nome string) *Escola {
	return &Escola{nome: nome, professores: []Professor{}}
}

func (e *Escola) AdicionarProfessor(professor *Professor) {
	for _, p := range e.professores {
		if p.nome == professor.nome {
			return
		}
	}
	e.professores = append(e.professores, *professor)
	professor.AdicionarEscola(e)
}

func (e *Escola) ListarProfessores() []string {
	nomes := []string{}
	for _, professor := range e.professores {
		nomes = append(nomes, professor.nome)
	}
	return nomes
}

func main() {
	professor1 := NewProfessor("Dr. Ricardo")
	professor2 := NewProfessor("Prof. Tatyana")

	escola1 := NewEscola("Escola Anglo")
	escola2 := NewEscola("Escola Marista")

	escola1.AdicionarProfessor(professor1)
	escola1.AdicionarProfessor(professor2)
	escola2.AdicionarProfessor(professor2)

	fmt.Printf("Professores na %s: %v\n", escola1.nome, escola1.ListarProfessores())
	fmt.Printf("Professores na %s: %v\n", escola2.nome, escola2.ListarProfessores())
	fmt.Printf("Escolas onde %s leciona: %v\n", professor1.nome, professor1.ListarEscolas())
	fmt.Printf("Escolas onde %s leciona: %v\n", professor2.nome, professor2.ListarEscolas())
}
