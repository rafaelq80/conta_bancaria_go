package model

import (
	"fmt"

	"github.com/fatih/color"
)

// Classe ContaCorrente
type ContaCorrente struct {
	Conta  // Composição
	limite float64
}

// Construtor para criar uma nova ContaCorrente
func NewContaCorrente(numero, agencia, tipo int, titular string, saldo, limite float64) *ContaCorrente {
	return &ContaCorrente{
		Conta: Conta{
			numero:  numero,
			agencia: agencia,
			tipo:    tipo,
			titular: titular,
			saldo:   saldo,
		},
		limite: limite,
	}
}

// Sobrescrita do método Sacar para ContaCorrente
func (cc *ContaCorrente) Sacar(valor float64) bool {

	saldoTotal := cc.saldo + cc.limite

	if valor > saldoTotal {

		color.Set(color.FgRed)
		fmt.Println("\nSaldo insuficiente!")
		color.Unset()

		return false
	}

	cc.saldo -= valor
	return true
}

// Método para obter o limite da conta corrente
func (cc *ContaCorrente) GetLimite() float64 {
	return cc.limite
}

// Método para definir o limite da conta corrente
func (cc *ContaCorrente) SetLimite(novoLimite float64) {
	cc.limite = novoLimite
}

// Sobrescrita do método Visualizar para ContaCorrente
func (cc *ContaCorrente) Visualizar() {
	cc.Conta.Visualizar() // Chama o método da conta base
	fmt.Printf("Limite da Conta: R$ %.2f \n", cc.limite)
}
