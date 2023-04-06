package main

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaMemory(a, b *int) int {
	*a = 50
	return *a + *b
}


func main() {
	//Cópias de variáveis
	myVar1 := 10
	myVar2 := 20

	/*Ao passar o myVar1, estamos mandando uma cópia de uma variável e 
	ela está sendo alterada de 10 para 50 na função soma*/
	println(soma(myVar1 , myVar2))
	//O valor continuará 10, porque somente a cópia foi alterada - Uma variável imutável
	println(myVar1)

	//Passando o endereço da memória
	println(somaMemory(&myVar1 , &myVar2))
	//O valor será 50, porque o valor foi alterado no endereço de memória do myVar1 - Uma variável mutável
	println(myVar1)
}