# Projeto Conta Bancária - Golang

<br />

<div align="center">
    <img src="https://i.imgur.com/JHinCnY.png" title="source: imgur.com" width="15%"/> 
    <img src="https://i.imgur.com/YC6Av6e.png" title="source: imgur.com" /> 
</div>

<br /><br />

## Diagramas de Classe

### Diagrama 01 - Classes Model

```mermaid
classDiagram
class Conta {
<<Abstract>>
  - Numero : int
  - Agencia : int
  - Tipo : int
  - Titular : string
  - Saldo : float64
  + GetNumero() int
  + GetAgencia() int
  + GetTipo() int
  + GetTitular() string
  + GetSaldo() float64
  + SetNumero(numero: int) void
  + SetAgencia(agencia: int) void
  + SetTipo(tipo: int) void
  + SetTitular(titular: string) void
  + SetSaldo(saldo: float64) void
  + Sacar(valor: float64) boolean
  + Depositar(valor: float64) void
  + Visualizar() void
}
class ContaCorrente {
  - Limite : float64
  + GetLimite() float64
  + SetLimite(limite: float64) void
  + Sacar(number float64) boolean
  + Visualizar() void
}
class ContaPoupanca {
  - Aniversario : int
  + GetAniversario() int
  + SetAniversario(aniversario: int) void
  + Visualizar() void
}

ContaCorrente *-- Conta : contém
ContaPoupanca *-- Conta : contém

```

<br />

### Diagrama 02 - Diagrama Completo


```mermaid
classDiagram
class Conta {
<<Abstract>>
  - Numero : int
  - Agencia : int
  - Tipo : int
  - Titular : string
  - Saldo : float64
  + GetNumero() int
  + GetAgencia() int
  + GetTipo() int
  + GetTitular() string
  + GetSaldo() float64
  + SetNumero(numero: int) void
  + SetAgencia(agencia: int) void
  + SetTipo(tipo: int) void
  + SetTitular(titular: string) void
  + SetSaldo(saldo: float64) void
  + Sacar(valor: float64) boolean
  + Depositar(valor: float64) void
  + Visualizar() void
}
class ContaCorrente {
  - Limite : float64
  + GetLimite() float64
  + SetLimite(limite: float64) void
  + Sacar(number float64) boolean
  + Visualizar() void
}
class ContaPoupanca {
  - Aniversario : int
  + GetAniversario() int
  + SetAniversario(aniversario: int) void
  + Visualizar() void
}
class IContaRepository{
<< Interface >>
+ BuscarPorNumero(numero: int) (IConta, error)
+ ListarTodas() ([]IConta, error)
+ Criar(conta: IConta) error
+ Atualizar(conta: IConta) error
+ Deletar(numero: int) error
+ Sacar(numero: int, valor: float64) error
+ Depositar(numero: int, valor: float64) error
+ Transferir(numeroOrigem: int, numeroDestino: int, valor: float64) error
}
class IConta{
<< Interface >>
+ GetNumero() int
+ SetNumero(int)
+ Sacar(float64) bool
+ Depositar(float64)
}
class ContaController{
-contas []IConta
-numeroAtual int
+NewContaController() IContaRepository
+ BuscarPorNumero(numero: int) (IConta, error)
+ ListarTodas() ([]IConta, error)
+ Criar(conta: IConta) error
+ Atualizar(conta: IConta) error
+ Deletar(numero: int) error
+ Sacar(numero: int, valor: float64) error
+ Depositar(numero: int, valor: float64) error
+ Transferir(numeroOrigem: int, numeroDestino: int, valor: float64) error
+ buscarConta(numero: int) (IConta, error)
}
Conta <--* ContaCorrente : contém
Conta <--* ContaPoupanca : contém
ContaController ..|> IContaRepository : implementa
ContaController --> IConta : usa
ContaCorrente ..|> IConta : implementa
ContaPoupanca ..|> IConta : implementa
IContaRepository ..> Conta : depende

```

<br /><br />

## Print da Tela

<div align="center">
   <img src="https://i.imgur.com/xZMR67v.png" title="source: imgur.com" />
</div>

<br /><br />

## Bibliotecas

- **Color** (https://github.com/fatih/color)
- **Survey** (https://github.com/AlecAivazis/survey/)

