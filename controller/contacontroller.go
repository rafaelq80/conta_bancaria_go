package controller

import (
	"errors"

	"github.com/rafaelq80/conta_bancaria_go/model"
	"github.com/rafaelq80/conta_bancaria_go/repository"
)

type ContaController struct {
	contas []interface{}
	numeroAtual int
}

func NewContaController() repository.IContaRepository {
	return &ContaController{
		contas: make([]interface{}, 0),
		numeroAtual: 0,
	}

}

func (cc *ContaController) Criar(conta interface{}) error {
	if !isValidAccountType(conta) {
		return errors.New("tipo de conta inválido")
	}

	cc.numeroAtual++

	if setter, ok := conta.(interface{ SetNumero(int) }); ok {
		setter.SetNumero(cc.numeroAtual)
	} else {
		return errors.New("a conta não possui um método SetNumero")
	}

	cc.contas = append(cc.contas, conta)
	return nil
}

func (cc *ContaController) BuscarPorNumero(numero int) (interface{}, error) {
	for _, conta := range cc.contas {
		if getNumero(conta) == numero {
			return conta, nil
		}
	}
	return nil, errors.New("conta não encontrada")
}

func (cc *ContaController) Atualizar(conta interface{}) error {
	if !isValidAccountType(conta) {
		return errors.New("tipo de conta inválido")
	}
	numero := getNumero(conta)
	for i, c := range cc.contas {
		if getNumero(c) == numero {
			cc.contas[i] = conta
			return nil
		}
	}
	return errors.New("conta não encontrada")
}

func (cc *ContaController) Deletar(numero int) error {

	for indice, conta := range cc.contas {
		if getNumero(conta) == numero {
			cc.contas = append(cc.contas[:indice], cc.contas[indice+1:]...)
			return nil
		}
	}
	
	return errors.New("conta não encontrada")
}

func (cc *ContaController) ListarTodas() ([]interface{}, error) {
	return cc.contas, nil
}

// Funções auxiliares

func getNumero(conta interface{}) int {

	switch c := conta.(type) {
		case *model.ContaCorrente:
			return c.GetNumero()
		case *model.ContaPoupanca:
			return c.GetNumero()
		default:
			panic("tipo de conta inválido")
	}
}

func isValidAccountType(conta interface{}) bool {
	switch conta.(type) {
	case model.ContaCorrente, *model.ContaCorrente, model.ContaPoupanca, *model.ContaPoupanca:
		return true
	default:
		return false
	}
}
