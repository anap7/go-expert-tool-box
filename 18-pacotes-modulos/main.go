package main

import (
	"fmt"
	"modulo/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado: ", s)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)

}

