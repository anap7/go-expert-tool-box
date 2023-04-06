package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(3, 5))
	fmt.Println(sum(35, 25))
	fmt.Println(sum2(35, 25, 4, 5, 9, 0, 78, 34, 143, 567, 257, 90))
}

//Funções comuns
func sum(a, b int) (int, error) {
	if a + b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}

//Funções variadicas - Quando você não sabe a quantidade de parametros a ser utilizado
func sum2(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}