package controller

import (
	"errors"
	"fmt"

	"github.com/rafaelq80/conta_bancaria_go/repository"
)

// ContaController implementa a Interface IContaRepository
type ContaController struct {
	contas      []Conta
	numeroAtual int
}

// Conta interface - Simula Polimorfismo
type Conta interface {
	GetNumero() int
	SetNumero(int)
	Sacar(float64) bool
	Depositar(float64)
}

// Método Cosbtrutor ContaController
func NewContaController() repository.IContaRepository {
	return &ContaController{
		contas:      make([]Conta, 0),
		numeroAtual: 0,
	}
}

func (cc *ContaController) Criar(conta interface{}) error {
	c, ok := conta.(Conta)
	if !ok {
		return errors.New("tipo de conta inválido")
	}

	cc.numeroAtual++
	c.SetNumero(cc.numeroAtual)
	cc.contas = append(cc.contas, c)
	return nil
}

func (cc *ContaController) ListarTodas() ([]interface{}, error) {
	listaContas := make([]interface{}, len(cc.contas))
	for i, c := range cc.contas {
		listaContas[i] = c
	}
	return listaContas, nil
}

func (cc *ContaController) BuscarPorNumero(numero int) (interface{}, error) {
	for _, conta := range cc.contas {
		if conta.GetNumero() == numero {
			return conta, nil
		}
	}
	return nil, fmt.Errorf("conta não encontrada: %d", numero)
}

func (cc *ContaController) Atualizar(conta interface{}) error {
	c, ok := conta.(Conta)
	if !ok {
		return errors.New("tipo de conta inválido")
	}

	for i, existingConta := range cc.contas {
		if existingConta.GetNumero() == c.GetNumero() {
			cc.contas[i] = c
			return nil
		}
	}
	return fmt.Errorf("conta não encontrada: %d", c.GetNumero())
}

func (cc *ContaController) Deletar(numero int) error {
	for i, conta := range cc.contas {
		if conta.GetNumero() == numero {
			cc.contas = append(cc.contas[:i], cc.contas[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("conta não encontrada: %d", numero)
}

func (cc *ContaController) Sacar(numero int, valor float64) error {
	conta, err := cc.buscarConta(numero)
	if err != nil {
		return err
	}

	if !conta.Sacar(valor) {
		return errors.New("saldo insuficiente")
	}
	return nil
}

func (cc *ContaController) Depositar(numero int, valor float64) error {
	conta, err := cc.buscarConta(numero)
	if err != nil {
		return err
	}

	conta.Depositar(valor)
	return nil
}

func (cc *ContaController) Transferir(numeroOrigem, numeroDestino int, valor float64) error {
	contaOrigem, err := cc.buscarConta(numeroOrigem)
	if err != nil {
		return fmt.Errorf("conta de origem não encontrada: %d", numeroOrigem)
	}

	contaDestino, err := cc.buscarConta(numeroDestino)
	if err != nil {
		return fmt.Errorf("conta de destino não encontrada: %d", numeroDestino)
	}

	if !contaOrigem.Sacar(valor) {
		return errors.New("saldo insuficiente na conta de origem")
	}

	contaDestino.Depositar(valor)
	return nil
}

// Método Auxiliar - Buscar Conta
func (cc *ContaController) buscarConta(numero int) (Conta, error) {
	for _, conta := range cc.contas {
		if conta.GetNumero() == numero {
			return conta, nil
		}
	}
	return nil, fmt.Errorf("conta não encontrada: %d", numero)
}