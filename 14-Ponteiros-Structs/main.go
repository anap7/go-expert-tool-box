package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c *Cliente) andou() {
	c.nome = "Dirce Maria"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func main() {
	user := Cliente{
		nome: "Dirce",
	}
	user.andou()
	fmt.Printf("O valor da struct com nome %v\n", user.nome)
	fmt.Print("\n")

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}