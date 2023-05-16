package main 

import "fmt"
//Declaração de constantes
const a = "Hello, World"

//Declaração da sua própria tipagem
type ID int

/*Ao declarar variáveis com tipos o go já insere um valor inicial */
var (
	b bool //Por padrão insere "false"
	c int //Por padrão insere 0
	d string //Por padrão é vazio
	e float64 //Por padrão o valor é +0.000000e+000
	f ID = 1 //Declarando o próprio tipo baseado no "type" do go
)

func main() {
	/*O Printf trabalha com o valor coringa %T que traz o tipo da variável*/
	fmt.Printf("O tipo da variável 'e' é %T", e)
	fmt.Print("\n")
	/*O Printf trabalha com o valor coringa %v que traz o valor da variável*/
	fmt.Printf("O tipo da variável 'e' é %v", e)

	/*
		%T = tipo da variável
		%v = valor da variável
	*/
}
