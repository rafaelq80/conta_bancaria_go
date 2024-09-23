package model

import (
    "fmt"
)

// Classe ContaPoupanca
type ContaPoupanca struct {
    Conta  // Composição
    aniversario int
}

// Construtor para criar uma nova ContaPoupanca
func NewContaPoupanca(numero, agencia, tipo int, titular string, saldo float64, aniversario int) *ContaPoupanca {
    return &ContaPoupanca{
        Conta: Conta{
            numero:  numero,
            agencia: agencia,
            tipo:    tipo,
            titular: titular,
            saldo:   saldo,
        },
        aniversario: aniversario,
    }
}

// Método para obter o aniversario da conta corrente
func (cc *ContaPoupanca) GetAniversario() int {
    return cc.aniversario
}

// Método para definir o aniversario da conta corrente
func (cc *ContaPoupanca) SetAniversario(aniversario int) {
    cc.aniversario = aniversario
}

// Sobrescrita do método Visualizar para ContaPoupanca
func (cc *ContaPoupanca) Visualizar() {
    cc.Conta.Visualizar() // Chama o método da conta base
    fmt.Printf("Aniversário da Conta: %d \n", cc.aniversario)
}