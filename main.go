package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/rafaelq80/conta_bancaria_go/controller"
	"github.com/rafaelq80/conta_bancaria_go/model"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var opcao = 0
	var numero, agencia, tipo, aniversario int
	var saldo, limite float64
	var titular string

	contaController := controller.NewContaController()

	contaController.Criar(model.NewContaCorrente(0, 123, 1, "João da Silva", 1000.00, 100.00))
	contaController.Criar(model.NewContaCorrente(0, 456, 1, "Maria dos Santos", 2900.00, 300.00))

	contaController.Criar(model.NewContaPoupanca(0, 789, 2, "Giovanna Benvenutti", 3200.00, 10))
	contaController.Criar(model.NewContaPoupanca(0, 123, 2, "Karina Girardini", 2000.00, 15))

	// c := model.NewConta(1, 123, 1, "João da Silva", 1000.00)
	// c.Visualizar()
	// c.Sacar(100.00)
	// fmt.Printf("\nSaldo: %.2f", c.GetSaldo())
	// c.Depositar(300.00)
	// fmt.Printf("\nSaldo: %.2f", c.GetSaldo())

	// cc := model.NewContaCorrente(2, 123, 1, "João da Silva", 1000.00, 500.00)
	// cc.Visualizar()
	// cc.Sacar(1100.00)
	// fmt.Printf("\nSaldo: %.2f", cc.GetSaldo())
	// cc.Depositar(300.00)
	// fmt.Printf("\nSaldo: %.2f", cc.GetSaldo())
	// cc.Sacar(800.00)
	// fmt.Printf("\nSaldo: %.2f", cc.GetSaldo())

	// cp := model.NewContaPoupanca(3, 123, 2, "João da Silva", 1000.00, 10)
	// cp.Visualizar()
	// cp.Sacar(1100.00)
	// fmt.Printf("\nSaldo: %.2f", cp.GetSaldo())
	// cp.Depositar(300.00)
	// fmt.Printf("\nSaldo: %.2f", cp.GetSaldo())

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
			color.Set(color.FgHiGreen)
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

			fmt.Println("Digite o número da agência: ")
			fmt.Scanln(&agencia)

			fmt.Println("Digite o tipo da conta (1-CC | 2-CP): ")
			fmt.Scanln(&tipo)

			fmt.Println("Digite o nome do titular da conta: ")
			titular, _ = reader.ReadString('\n')

			fmt.Println("Digite o saldo da conta: ")
			fmt.Scanln(&saldo)

			switch (tipo) {
				case 1:

					fmt.Println("Digite o limite da conta: ")
					fmt.Scanln(&limite)
			
					resposta := contaController.Criar(model.NewContaCorrente(0, agencia, tipo, titular, saldo, limite))

					if resposta == nil {
						fmt.Println("Conta Corrente criada com sucesso!")
					}

				case 2:

					fmt.Println("Digite o dia do aniversário da conta: ")
					fmt.Scanln(&aniversario)

					contaController.Criar(model.NewContaPoupanca(0, agencia, tipo, titular, saldo, aniversario))

					fmt.Println("Conta Poupança criada com sucesso!")

			}


			keyPress()
		case 2:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nListar todas as Contas\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO LISTAR TODAS AS CONTAS */

			contas, err := contaController.ListarTodas()

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			for _, conta := range contas {
				switch c := conta.(type) {
				case *model.ContaCorrente:
					c.Visualizar()
				case *model.ContaPoupanca:
					c.Visualizar()
				default:
					fmt.Println("\nTipo de Conta Inválida")
				}
			}

			keyPress()
		case 3:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nConsultar dados da Conta - por número\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO CONSULTAR CONTA PELO NÚMERO */

			fmt.Println("Digite o número da conta: ")
			fmt.Scanln(&numero)

			conta, err := contaController.BuscarPorNumero(numero)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			switch c := conta.(type) {
			case *model.ContaCorrente:
				c.Visualizar()
			case *model.ContaPoupanca:
				c.Visualizar()
			default:
				fmt.Println("\nTipo de Conta Inválida")
			}

			keyPress()
		case 4:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nAtualizar dados da Conta\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO ATUALIZAR DADOS DE UMA CONTA */

			keyPress()
		case 5:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nApagar uma Conta\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO EXCLUIR CONTA */

			keyPress()
		case 6:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nSaque\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO SACAR */

			keyPress()
		case 7:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nDepósito\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO DEPOSITAR */

			keyPress()
		case 8:
			color.Set(color.FgYellow, color.Bold)
			fmt.Print("\n\nTransferência entre Contas\n\n")
			color.Unset()

			/* CHAMADA DO MÉTODO TRANSFERIR */

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
