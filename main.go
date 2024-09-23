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
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var opcao = 0
	var numero, agencia, tipo, aniversario, numeroDestino int
	var saldo, limite, valor float64
	var titular string

	contaController := controller.NewContaController()

	contaController.Criar(model.NewContaCorrente(0, 123, 1, "João da Silva", 1000.00, 100.00))
	contaController.Criar(model.NewContaCorrente(0, 456, 1, "Maria dos Santos", 2900.00, 300.00))

	contaController.Criar(model.NewContaPoupanca(0, 789, 2, "Giovanna Benvenutti", 3200.00, 10))
	contaController.Criar(model.NewContaPoupanca(0, 123, 2, "Karina Girardini", 2000.00, 15))

	for true {

		color.Set(color.BgBlack, color.FgYellow)

		fmt.Println("*****************************************************")
		fmt.Println("                                                     ")
		fmt.Println("                BANCO DO BRAZIL COM Z                ")
		fmt.Println("                                                     ")
		fmt.Println("*****************************************************")
		fmt.Println("                                                     ")
		fmt.Println("            1 - Criar Conta                          ")
		fmt.Println("            2 - Listar todas as Contas               ")
		fmt.Println("            3 - Buscar Conta por Numero              ")
		fmt.Println("            4 - Atualizar Dados da Conta             ")
		fmt.Println("            5 - Apagar Conta                         ")
		fmt.Println("            6 - Sacar                                ")
		fmt.Println("            7 - Depositar                            ")
		fmt.Println("            8 - Transferir valores entre Contas      ")
		fmt.Println("            9 - Sair                                 ")
		fmt.Println("                                                     ")
		fmt.Println("*****************************************************")
		fmt.Println("                                                     ")

		fmt.Println("Digite a opção desejada: ")

		color.Unset()

		fmt.Scanln(&opcao)

		if opcao == 9 {
			color.Set(color.FgCyan)
			fmt.Println("\nBanco do Brazil com Z - O seu Futuro começa aqui!")
			sobre()
			os.Exit(0)
			color.Unset()
		}

		switch opcao {
		case 1:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nCriar Conta\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO CRIAR CONTA */

			fmt.Println("\nDigite o número da agência: ")
			fmt.Scanln(&agencia)

			fmt.Print("\n")

			selecionada := ""
			prompt := &survey.Select{
				Message: "Selecione o tipo da conta:",
				Options: []string{"Conta Corrente", "Conta Poupança"},
			}
			survey.AskOne(prompt, &selecionada)

			switch selecionada {
			case "Conta Corrente":
				tipo = 1
			case "Conta Poupança":
				tipo = 2
			}

			fmt.Println("\nDigite o nome do titular da conta: ")
			titular, _ = reader.ReadString('\n')
			titular = strings.TrimSpace(titular)

			fmt.Println("\nDigite o saldo da conta: ")
			fmt.Scanln(&saldo)

			var resposta error

			switch tipo {
			case 1:

				fmt.Println("\nDigite o limite da conta: ")
				fmt.Scanln(&limite)

				resposta = contaController.Criar(model.NewContaCorrente(0, agencia, tipo, titular, saldo, limite))

			case 2:

				fmt.Println("\nDigite o dia do aniversário da conta: ")
				fmt.Scanln(&aniversario)

				resposta = contaController.Criar(model.NewContaPoupanca(0, agencia, tipo, titular, saldo, aniversario))

			}

			if resposta == nil {
				color.Set(color.FgGreen)
				fmt.Printf("\nConta %s criada com sucesso!\n", selecionada)
				color.Unset()
			} else {
				color.Set(color.FgRed)
				fmt.Printf("\nErro ao criar conta: %s\n", resposta)
				color.Unset()
			}

			keyPress()
		case 2:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nListar todas as Contas\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO LISTAR TODAS AS CONTAS */

			contas, err := contaController.ListarTodas()

			if err != nil {
				color.Set(color.FgRed)
				fmt.Println("Error:", err)
				color.Unset()
				return
			}

			for _, conta := range contas {
				switch c := conta.(type) {
				case *model.ContaCorrente:
					c.Visualizar()
				case *model.ContaPoupanca:
					c.Visualizar()
				default:
					color.Set(color.FgRed)
					fmt.Println("\nTipo de Conta Inválida")
					color.Unset()
				}
			}

			color.Set(color.FgCyan)
			fmt.Printf("\nTotal de Contas Cadastradas: %d", len(contas))
			color.Unset()

			keyPress()
		case 3:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nConsultar dados da Conta - por número\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO CONSULTAR CONTA PELO NÚMERO */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			conta, err := contaController.BuscarPorNumero(numero)

			if err == nil {
				switch c := conta.(type) {
				case *model.ContaCorrente:
					c.Visualizar()
				case *model.ContaPoupanca:
					c.Visualizar()
				default:
					color.Set(color.FgRed)
					fmt.Println("\nTipo de Conta Inválida")
					color.Unset()
				}
			} else {

				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()

			}

			keyPress()
		case 4:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nAtualizar dados da Conta\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO ATUALIZAR DADOS DE UMA CONTA */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			conta, err := contaController.BuscarPorNumero(numero)

			if err == nil {

				switch c := conta.(type) {
				case *model.ContaCorrente:
					tipo = c.GetTipo()
				case *model.ContaPoupanca:
					tipo = c.GetTipo()
				default:
					color.Set(color.FgRed)
					fmt.Println("\nTipo de Conta Inválida")
					color.Unset()
				}

				fmt.Println("\nDigite o número da agência: ")
				fmt.Scanln(&agencia)

				fmt.Println("\nDigite o nome do titular da conta: ")
				titular, _ = reader.ReadString('\n')
				titular = strings.TrimSpace(titular)

				fmt.Println("\nDigite o saldo da conta: ")
				fmt.Scanln(&saldo)

				var resposta error

				switch tipo {
				case 1:

					fmt.Println("\nDigite o limite da conta: ")
					fmt.Scanln(&limite)

					resposta = contaController.Atualizar(model.NewContaCorrente(numero, agencia, tipo, titular, saldo, limite))

				case 2:

					fmt.Println("\nDigite o dia do aniversário da conta: ")
					fmt.Scanln(&aniversario)

					resposta = contaController.Atualizar(model.NewContaPoupanca(numero, agencia, tipo, titular, saldo, aniversario))

				}

				if resposta == nil {
					color.Set(color.FgGreen)
					fmt.Printf("Conta Número %d atualizada com sucesso!\n", numero)
					color.Unset()
				} else {
					color.Set(color.FgRed)
					fmt.Printf("\n%s\n", resposta)
					color.Unset()
				}

			} else {

				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()

			}

			keyPress()
		case 5:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nApagar uma Conta\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO EXCLUIR CONTA */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			err := contaController.Deletar(numero)

			if err != nil {
				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()
			} else {
				color.Set(color.FgGreen)
				fmt.Printf("\n\nConta Número %d excluída com sucesso!", numero)
				color.Unset()
			}

			keyPress()
		case 6:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nSaque\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO SACAR */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			fmt.Println("Digite o valor do saque (R$): ")
			fmt.Scanln(&valor)

			err := contaController.Sacar(numero, valor)

			if err != nil {
				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()
			}else{

				color.Set(color.FgGreen)
				fmt.Printf("\nO Saque no valor de R$ %.2f, foi realizado com sucesso!", valor)
				color.Unset()

			}

			keyPress()
		case 7:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nDepósito\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO DEPOSITAR */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			fmt.Println("Digite o valor do depósito (R$): ")
			fmt.Scanln(&valor)

			err := contaController.Depositar(numero, valor)

			if err != nil {
				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()
			} else {
				color.Set(color.FgGreen)
				fmt.Printf("\n\nO depósito de R$ %.2f, na conta Número %d, foi efetuado com sucesso!", valor, numero)
				color.Unset()
			}

			keyPress()
		case 8:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nTransferência entre Contas\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO TRANSFERIR */

			fmt.Println("Digite o número da conta de origem: ")
			fmt.Scanln(&numero)

			fmt.Println("Digite o número da conta de destino: ")
			fmt.Scanln(&numeroDestino)

			fmt.Println("Digite o valor do saque (R$): ")
			fmt.Scanln(&valor)

			err := contaController.Transferir(numero, numeroDestino, valor)

			if err != nil {
				color.Set(color.FgRed)
				fmt.Printf("\n%s", err)
				color.Unset()
			} else {
				color.Set(color.FgGreen)
				fmt.Printf("\n\nA Transferência no valor de R$ %.2f, para a conta Número %d, foi efetuada com sucesso!", valor, numeroDestino)
				color.Unset()
			}

			keyPress()
		default:
			color.Set(color.FgRed, color.Bold)
			fmt.Print("\nOpção Inválida!\n")
			color.Unset()

			keyPress()
		}

	}
}

func sobre() {
	fmt.Println("\n*****************************************************")
	fmt.Println("Projeto Desenvolvido por: ")
	fmt.Println("Generation Brasil - generation@generation.org")
	fmt.Println("github.com/conteudoGeneration")
	fmt.Println("*****************************************************")
}

func keyPress() {
	fmt.Println("\n\nPressione Enter para continuar...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
