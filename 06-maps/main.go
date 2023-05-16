package main

import "fmt"

func main() {
	salaries := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salaries["Wesley"])
	delete (salaries, "Wesley")
	salaries["Claudinho"] = 5000
	fmt.Println(salaries)
	fmt.Println(salaries["Claudinho"])

	//Iniciando um map
	/*Aqui você está declarando um slice cuja chave é uma string e o valor é int*/
	sal := make(map[string]int)

	//Opção 2 
	sal1:= map[string]int{}

	//Adicionando novos itens
	sal["Novo"] = 200
	sal["Novo"] = 200

	sal1["New"] = 300

	fmt.Println(sal)
	fmt.Println(sal1)

	for name, salaries := range salaries {
		fmt.Printf("O salario de %s é %d\n", name, salaries)
	}

	for _, salaries := range salaries {
		fmt.Printf("O salario é %d\n", salaries)
	}
}