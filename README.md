# Projeto Conta Bancária - Golang

<br />

<div align="center">
    <img src="https://i.imgur.com/JHinCnY.png" title="source: imgur.com" width="15%"/> 
    <img src="https://i.imgur.com/YC6Av6e.png" title="source: imgur.com" /> 
</div>
<br /><br />

Este projeto implementa um sistema básico de gerenciamento de contas bancárias utilizando a linguagem Go, utilizando os conceitos da Programação Orientada a Objetos oferecidos pela Linguagem. O sistema oferece funcionalidades para criar, gerenciar e realizar operações bancárias básicas em diferentes tipos de contas bancárias.

<br />

## Principais Características

1. **Gerenciamento de Contas**: 
   - Criação de diferentes tipos de contas (Conta Corrente e Conta Poupança)
   - Listagem de todas as contas
   - Busca de conta por número
   - Atualização e remoção de contas

2. **Operações Bancárias**: 
   - Saque
   - Depósito
   - Transferência entre contas

3. **Polimorfismo**: 
   - Utilização de interface `IConta` para permitir operações uniformes em diferentes tipos de contas, simulando o Polimorfismo em Golang

4. **Controle de Números de Conta**: 
   - Geração automática de números de conta, simulando a Chave Primária de um banco de dados

## Estrutura do Projeto

O projeto é estruturado em diferentes componentes:

1. **Controller (ContaController)**:
   - Implementa a lógica de negócios
   - Gerencia uma coleção de contas
   - Implementa a interface `IContaRepository`

2. **Interfaces**:
   - `IConta`: Define os métodos básicos que todos os tipos de conta devem implementar
   - `IContaRepository`: Define as operações de CRUD e as transações bancárias

3. **Modelos de Conta**:
   - Implementação do Modelo de dados da aplicação (`Conta`, `ContaCorrente` e `ContaPoupanca`)
   - Implementações concretas da Interface `IConta` 

## Tecnologias

- **Linguagem**: Go
- **Princípios de Design**: 
  - Programação Orientada a Objetos
  - Relacionamento entre Classes do tipo Associação por Composição
  - Inversão de Dependência (Injeção de Dependências)
  - Simulação de Polimorfismo através de Interfaces

## Funcionalidades Detalhadas

- Criação de contas com número único gerado automaticamente
- Listagem de todas as contas cadastradas
- Busca de conta específica por número
- Atualização de informações de conta
- Remoção de contas do sistema
- Operações de saque com verificação de saldo
- Depósitos em contas
- Transferências entre contas com validação

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
+ GetTitular() string
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
+ BuscarPorTitular(titular string) ([]IConta, error)
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

<br /><br />

## Executar o Projeto



Para executar um projeto em Go, siga estes passos básicos:

1. Instale o Go na sua máquina, caso ele ainda não esteja instalado (https://go.dev/)

2. Clone o projeto na sua máquina

3. Abra o projeto no Visual Studio Code

4. Abra o Terminal do Visual Studio Code

5. Instale todas as Dependências do projeto, através do comando abaixo:

```bash
go mod tidy
```

6. As dependências serão registradas automaticamente no seu arquivo `go.mod`.

7. Execute o Projeto, via terminal, através do comando abaixo:


```bash
go run main.go
```

8. Para compilar o seu projeto em um arquivo executável (exe), use o comando abaixo:

```bash
go build main.go
```

9. O comando acima gerará um arquivo executável chamado **main.exe**. 
10. Você pode executar o arquivo gerado através do comando abaixo:

```bash
./main.exe  # Para Unix/Linux/Mac
main.exe  # Para Windows
```

