class ContaBancaria {
    private double saldo;
    public String titular;

    public ContaBancaria(String titular, double saldoInicial) {
        this.saldo = saldoInicial;
        this.titular = titular;
    }

    public String depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            return String.format("Depósito de R$%.2f realizado. Saldo atual: R$%.2f", valor, saldo);
        } else {
            return "Valor de depósito inválido!";
        }
    }

    public String sacar(double valor) {
        if (valor > 0 && valor <= saldo) {
            saldo -= valor;
            return String.format("Saque de R$%.2f realizado. Saldo atual: R$%.2f", valor, saldo);
        } else {
            return "Saldo insuficiente ou valor de saque inválido!";
        }
    }

    public String exibirSaldo() {
        return String.format("Saldo atual de %s: R$%.2f", titular, saldo);
    }
}

public class Main {
    public static void main(String[] args) {
        ContaBancaria contaTeste = new ContaBancaria("João Silva", 1000);
        System.out.println(contaTeste.depositar(500));
        System.out.println(contaTeste.sacar(300));
        System.out.println(contaTeste.exibirSaldo());
    }
}
