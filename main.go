package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/rafaelq80/conta_bancaria_go/controller"
	"github.com/rafaelq80/conta_bancaria_go/model"
	"github.com/rafaelq80/conta_bancaria_go/repository"
)

// Banco interface - Simula Polimorfismo
type Banco interface {
    Visualizar()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	contaController := controller.NewContaController()
	
	inicializaContas(contaController)

	for {
		
		exibirMenu()

		opcao := readOption()

		switch opcao {
		case 1:
			criarConta(contaController, reader)
		case 2:
			listarContas(contaController)
		case 3:
			consultarConta(contaController)
		case 4:
			atualizarConta(contaController, reader)
		case 5:
			apagarConta(contaController)
		case 6:
			realizarSaque(contaController)
		case 7:
			realizarDeposito(contaController)
		case 8:
			realizarTransferencia(contaController)
		case 9:
			encerrarPrograma()
		default:
			color.Set(color.FgRed, color.Bold)
			fmt.Print("\nOpção Inválida!\n")
			color.Unset()
		}

		keyPress()
	}
}

func inicializaContas(contaController repository.IContaRepository) {
	contaController.Criar(model.NewContaCorrente(0, 123, 1, "João da Silva", 1000.00, 100.00))
	contaController.Criar(model.NewContaCorrente(0, 456, 1, "Maria dos Santos", 2900.00, 300.00))
	contaController.Criar(model.NewContaPoupanca(0, 789, 2, "Giovanna Benvenutti", 3200.00, 10))
	contaController.Criar(model.NewContaPoupanca(0, 123, 2, "Karina Girardini", 2000.00, 15))
}

func exibirMenu() {

	color.Set(color.BgBlack, color.FgYellow)
	fmt.Println("\n*****************************************************")
	fmt.Println("                BANCO DO BRAZIL COM Z                ")
	fmt.Println("*****************************************************")
	fmt.Println("            1 - Criar Conta                          ")
	fmt.Println("            2 - Listar todas as Contas               ")
	fmt.Println("            3 - Buscar Conta por Numero              ")
	fmt.Println("            4 - Atualizar Dados da Conta             ")
	fmt.Println("            5 - Apagar Conta                         ")
	fmt.Println("            6 - Sacar                                ")
	fmt.Println("            7 - Depositar                            ")
	fmt.Println("            8 - Transferir valores entre Contas      ")
	fmt.Println("            9 - Sair                                 ")
	fmt.Println("*****************************************************")
	color.Unset()
}

func readOption() int {
	var opcao int
	fmt.Print("\nDigite a opção desejada: ")
	fmt.Scanln(&opcao)
	return opcao
}

func criarConta(contaController repository.IContaRepository, reader *bufio.Reader) {
	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nCriar Conta\n\n")
	color.Unset()

	agencia := readInt("\nDigite o número da agência: ")
	tipo := selectTipoConta()
	titular := readString(reader, "\nDigite o nome do titular da conta: ")
	saldo := readFloat("\nDigite o saldo da conta: ")

	var resposta error

	if tipo == 1 {
		limite := readFloat("Digite o limite da conta: ")
		resposta = contaController.Criar(model.NewContaCorrente(0, agencia, tipo, titular, saldo, limite))
	} else {
		aniversario := readInt("Digite o dia do aniversário da conta: ")
		resposta = contaController.Criar(model.NewContaPoupanca(0, agencia, tipo, titular, saldo, aniversario))
	}

	if resposta == nil {
		color.Green("\nConta criada com sucesso!")
	} else {
		color.Red("\nErro ao criar conta: %s", resposta)
	}
}

func listarContas(contaController repository.IContaRepository) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nListar todas as Contas\n\n")
	color.Unset()

	contas, err := contaController.ListarTodas()

	if err != nil {
		color.Red("Error: %s", err)
		return
	}

	for _, conta := range contas {
        if c, ok := conta.(Banco); ok {
            c.Visualizar()
        } else {
            color.Red("\nTipo de Conta Inválida")
        }
    }

	color.Cyan("\nTotal de Contas Cadastradas: %d", len(contas))
}

func consultarConta(contaController repository.IContaRepository) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nConsultar dados da Conta - por número\n\n")
	color.Unset()

	numero := readInt("Digite o número da conta: ")
	conta, err := contaController.BuscarPorNumero(numero)

	if err == nil {
		if c, ok := conta.(Banco); ok {
            c.Visualizar()
        } else {
            color.Red("\nTipo de Conta Inválida")
        }
	} else {
		color.Red("\n%s", err)
	}
}

func atualizarConta(contaController repository.IContaRepository, reader *bufio.Reader) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nAtualizar dados da Conta\n\n")
	color.Unset()

	numero := readInt("Digite o número da conta: ")
	conta, err := contaController.BuscarPorNumero(numero)

	if err == nil {
		agencia := readInt("\nDigite o número da agência: ")
		titular := readString(reader, "\nDigite o nome do titular da conta: ")
		saldo := readFloat("\nDigite o saldo da conta: ")

		var resposta error
		switch c := conta.(type) {
			case *model.ContaCorrente:
				limite := readFloat("\nDigite o limite da conta: ")
				resposta = contaController.Atualizar(model.NewContaCorrente(numero, agencia, c.GetTipo(), titular, saldo, limite))
			case *model.ContaPoupanca:
				aniversario := readInt("\nDigite o dia do aniversário da conta: ")
				resposta = contaController.Atualizar(model.NewContaPoupanca(numero, agencia, c.GetTipo(), titular, saldo, aniversario))
			default:
				color.Red("\nTipo de Conta Inválida")
				return
		}

		if resposta == nil {
			color.Green("Conta Número %d atualizada com sucesso!", numero)
		} else {
			color.Red("\n%s", resposta)
		}
	} else {
		color.Red("\n%s", err)
	}
}

func apagarConta(contaController repository.IContaRepository) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nApagar uma Conta\n\n")
	color.Unset()

	numero := readInt("Digite o número da conta: ")
	err := contaController.Deletar(numero)

	if err != nil {
		color.Red("\n%s", err)
	} else {
		color.Green("\nConta Número %d excluída com sucesso!", numero)
	}
}

func realizarSaque(contaController repository.IContaRepository) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nSaque\n\n")
	color.Unset()

	numero := readInt("Digite o número da conta: ")
	valor := readFloat("Digite o valor do saque (R$): ")

	err := contaController.Sacar(numero, valor)

	if err != nil {
		color.Red("\n%s", err)
	} else {
		color.Green("\nO Saque no valor de R$ %.2f, foi realizado com sucesso!", valor)
	}
}

func realizarDeposito(contaController repository.IContaRepository) {

	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nDepósito\n\n")
	color.Unset()

	numero := readInt("Digite o número da conta: ")
	valor := readFloat("Digite o valor do depósito (R$): ")

	err := contaController.Depositar(numero, valor)

	if err != nil {
		color.Red("\n%s", err)
	} else {
		color.Green("\nO depósito de R$ %.2f, na conta Número %d, foi efetuado com sucesso!", valor, numero)
	}
}

func realizarTransferencia(contaController repository.IContaRepository) {
	
	color.Set(color.FgYellow, color.Bold)
	fmt.Print("\n\nTransferência entre Contas\n\n")
	color.Unset()

	numeroOrigem := readInt("Digite o número da conta de origem: ")
	numeroDestino := readInt("Digite o número da conta de destino: ")
	valor := readFloat("Digite o valor da transferência (R$): ")

	err := contaController.Transferir(numeroOrigem, numeroDestino, valor)

	if err != nil {
		color.Red("\n%s", err)
	} else {
		color.Green("\nA Transferência no valor de R$ %.2f, para a conta Número %d, foi efetuada com sucesso!", valor, numeroDestino)
	}
}

func encerrarPrograma() {
	color.Cyan("\nBanco do Brazil com Z - O seu Futuro começa aqui!")
	sobre()
	os.Exit(0)
}

func sobre() {
	fmt.Println("\n*****************************************************")
	fmt.Println("Projeto Desenvolvido por: ")
	fmt.Println("Generation Brasil - generation@generation.org")
	fmt.Println("github.com/conteudoGeneration")
	fmt.Println("*****************************************************")
}

func keyPress() {
	fmt.Print("\nPressione Enter para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func readInt(prompt string) int {
	var value int
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

func readFloat(prompt string) float64 {
	var value float64
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

func readString(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func selectTipoConta() int {
	selecionada := ""
	prompt := &survey.Select{
		Message: "Selecione o tipo da conta:",
		Options: []string{"Conta Corrente", "Conta Poupança"},
	}
	survey.AskOne(prompt, &selecionada)

	switch selecionada {
	case "Conta Corrente":
		return 1
	case "Conta Poupança":
		return 2
	default:
		return 0
	}
}