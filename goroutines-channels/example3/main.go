package main

import "fmt"

//Funcao main eh uma thread (thread 01)
func main() {
	// Criar um canal do tipo string
	hello := make(chan string)

	// thread 02 - eh uma go routine anonima
	go func() {
		// O valor que estou passando para o canal eh "Hello world"
		hello <- "Hello World"
	}()

	//Voltando para a thread 01
	// Aqui estou falando que se eu tiver um valor nesse canal "hello", vc passa pra variavel result
	result := <-hello

	fmt.Println(result)
}