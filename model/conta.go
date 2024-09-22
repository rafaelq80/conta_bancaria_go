package model

import (
    "fmt"
)

// Classe base Conta
type Conta struct {
    numero  int
    agencia int
    tipo    int
    titular string
    saldo   float64
}

// Construtor para criar uma nova Conta
func NewConta(numero, agencia, tipo int, titular string, saldo float64) *Conta {
    return &Conta{
        numero:  numero,
        agencia: agencia,
        tipo:    tipo,
        titular: titular,
        saldo:   saldo,
    }
}

// Getters
func (c *Conta) GetNumero() int {
    return c.numero
}

func (c *Conta) GetAgencia() int {
    return c.agencia
}

func (c *Conta) GetTipo() int {
    return c.tipo
}

func (c *Conta) GetTitular() string {
    return c.titular
}

func (c *Conta) GetSaldo() float64 {
    return c.saldo
}

// Setters
func (c *Conta) SetNumero(numero int) {
    c.numero = numero
}

func (c *Conta) SetAgencia(agencia int) {
    c.agencia = agencia
}

func (c *Conta) SetTipo(tipo int) {
    c.tipo = tipo
}

func (c *Conta) SetTitular(titular string) {
    c.titular = titular
}

func (c *Conta) SetSaldo(saldo float64) {
    c.saldo = saldo
}

// Método para sacar
func (c *Conta) Sacar(valor float64) bool {
    if valor > c.saldo {
        fmt.Println("\nSaldo insuficiente!")
        return false
    }
    c.saldo -= valor
    fmt.Printf("\nSaque de R$ %.2f realizado com sucesso. \nNovo Saldo da Conta: R$ %.2f\n", valor, c.saldo)
    return true
}

// Método para depositar
func (c *Conta) Depositar(valor float64) {
    c.saldo += valor
}

func (c *Conta) Visualizar() {

    tipo := ""

    switch c.tipo {
    case 1:
        tipo = "Conta Corrente"
    case 2:
        tipo = "Conta Poupança"
    }

    fmt.Println("\n***********************************************")
    fmt.Println("                Dados da Conta                 ")
    fmt.Println("***********************************************")
    fmt.Printf("\nNúmero da Conta: %d", c.numero)
    fmt.Printf("\nNúmero da Agência: %d", c.agencia)
    fmt.Printf("\nTipo da Conta: %s", tipo)
    fmt.Printf("\nTitular da Conta: %s", c.titular)
    fmt.Printf("\nSaldo da Conta: %.2f \n", c.saldo)
}