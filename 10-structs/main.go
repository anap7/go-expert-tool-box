package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco // ou Endereco Endereco
}


func main() {
	user := Cliente{
		Nome:  "Ana",
		Idade: 30,
		Ativo: true,
	}

	//Mudando
	user.Idade = 24

	user.Cidade = "São Paulo"
	user.Logradouro = "Rua show"

	user.Endereco.Cidade = "Santa Catarina"

	user.Desativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", user.Nome, user.Idade, user.Ativo)

}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}