interface Imprimivel {
    void imprimir();
}

class Relatorio implements Imprimivel {
    @Override
    public void imprimir() {
        System.out.println("Imprimindo o Relatório.");
    }
}

class Contrato implements Imprimivel {
    @Override
    public void imprimir() {
        System.out.println("Imprimindo o Contrato.");
    }
}

public class Main {
    public static void main(String[] args) {
        Imprimivel relatorio = new Relatorio();
        Imprimivel contrato = new Contrato();

        relatorio.imprimir();
        contrato.imprimir();
    }
}
