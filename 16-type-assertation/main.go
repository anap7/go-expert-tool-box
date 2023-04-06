package main

import "fmt"

func main() {
	var minhaVar interface{} = "Dirce"

	println(minhaVar)
	println(minhaVar.(string))

	//Retorna 2 valores, um convertido para inteiro (res) e se deu certo (ok)
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)
}
