package main

func main() {
	/*Memoria -> Endereço -> Valor
	Quando a é igual 10, o go abre uma "caixinha" na memória que tem um endereço com 
	o valor 10. Quando alguem quiser saber o valor de a, você vai pelo endereço da memoria*/
	a := 10

	println(a) //Valor 
	println(&a) //Endereço na memoria

	/*Quando usamos o "*" estamos criando um ponteiro (endereçamento na memoria)
	A variável aponta para o ponteiro que tem um endereço na memória que possui um valor */
	var pointer *int = &a
	println(pointer)
	*pointer = 20
	//O valor de a será 20, porque o pointer aponta para o mesmo endereço de memoria do a
	println(a)
	//Ao criar uma variavel simples e atribuir o ponteiro, ele se transforma em ponteiro
	b := &a
	println(b)
	//E para saber o valor de b usamos o "*", porque vc pergunta para o ponteiro qual o valor desse endereço
	println(*b)
	*b = 30
	//Será 30, porque aponta pro mesmo endereço de memória
	println(a)
}