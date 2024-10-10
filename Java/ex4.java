abstract class Animal {
    protected String nome;

    public Animal(String nome) {
        this.nome = nome;
    }

    public abstract String som();
}

class Cachorro extends Animal {
    public Cachorro(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return nome + " faz: Au Au!";
    }
}

class Gato extends Animal {
    public Gato(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return nome + " faz: Miau!";
    }
}

public class Main {
    public static void main(String[] args) {
        Animal cachorro = new Cachorro("Thor");
        Animal gato = new Gato("Mel");

        System.out.println(cachorro.som());
        System.out.println(gato.som());
    }
}
