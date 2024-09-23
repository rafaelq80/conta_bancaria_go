package repository

type IContaRepository interface {
	
    Criar(conta interface{}) error
	BuscarPorNumero(numero int) (interface{}, error)
	Atualizar(conta interface{}) error
	Deletar(numero int) error
	ListarTodas() ([]interface{}, error)

}