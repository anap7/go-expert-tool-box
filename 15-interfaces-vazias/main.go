package main

import "fmt"

//Significa que a interface se aplica a todo mundo
//Uma interface vazia suporta qualquer coisa
type x interface {}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello Dirce"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}