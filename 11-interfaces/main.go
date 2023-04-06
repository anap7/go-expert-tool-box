package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

/*Qualquer struct que tiver o método desativar está implementando a interface pessoa, isso ocorre de maneira automaticamente*/
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco // ou Endereco Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	user := Cliente{
		Nome: "Ana",
		Idade: 30,
		Ativo: false,
	}

	user.Ativo = true
	user.Desativar()

	//Eu consigo passar o cliente para a desativacao, o cliente implementa Pessoa e isso faz com que o cliente seja pessoa
	Desativacao(user)
}

