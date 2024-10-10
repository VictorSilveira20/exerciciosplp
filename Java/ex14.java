class Configuracao {
    private static Configuracao instancia;
    private String configuracao;

    private Configuracao() {
        configuracao = "Configuração padrão";
    }

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    public void setConfiguracao(String configuracao) {
        this.configuracao = configuracao;
    }

    public String getConfiguracao() {
        return configuracao;
    }
}

public class Main {
    public static void main(String[] args) {
        Configuracao config1 = Configuracao.getInstancia();
        config1.setConfiguracao("Host: localhost");

        Configuracao config2 = Configuracao.getInstancia();
        config2.setConfiguracao("Port: 8080");

        System.out.println(config1.getConfiguracao());
        System.out.println(config2.getConfiguracao());
        System.out.println(config1 == config2);
    }
}
