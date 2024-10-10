import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    private List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this);
        }
    }

    public List<String> listarEscolas() {
        List<String> nomesEscolas = new ArrayList<>();
        for (Escola escola : escolas) {
            nomesEscolas.add(escola.getNome());
        }
        return nomesEscolas;
    }

    public String getNome() {
        return nome;
    }
}

class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
            professor.adicionarEscola(this);
        }
    }

    public List<String> listarProfessores() {
        List<String> nomesProfessores = new ArrayList<>();
        for (Professor professor : professores) {
            nomesProfessores.add(professor.getNome());
        }
        return nomesProfessores;
    }

    public String getNome() {
        return nome;
    }
}

public class Main {
    public static void main(String[] args) {
        Professor professor1 = new Professor("Dr. Ricardo");
        Professor professor2 = new Professor("Prof. Tatyana");

        Escola escola1 = new Escola("Escola Anglo");
        Escola escola2 = new Escola("Escola Marista");

        escola1.adicionarProfessor(professor1);
        escola1.adicionarProfessor(professor2);
        escola2.adicionarProfessor(professor2);

        System.out.println("Professores na " + escola1.getNome() + ": " + escola1.listarProfessores());
        System.out.println("Professores na " + escola2.getNome() + ": " + escola2.listarProfessores());
        System.out.println("Escolas onde " + professor1.getNome() + " leciona: " + professor1.listarEscolas());
        System.out.println("Escolas onde " + professor2.getNome() + " leciona: " + professor2.listarEscolas());
    }
}
