package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Criando um arquivo chamado file.txt
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	//Escrevendo no arquivo
	size, err := f.Write([]byte("Writing data on file"))
	//size, err := f.WriteString("Hello, Dirce")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File was created successfully! Size: %d bytes", size)
	f.Close()

	//Lendo um arquivo
	f2, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n" + string(f2))

	//leitura de pouco em pouco abrindo o arquivo
	//Abrindo o arquivo
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	//Lendo o arquivo
	reader := bufio.NewReader(file)
	buffer := make([]byte, 3)
	for {
		//Faz a leitura baseada no buffer que definimos acima, ou seja, somente 10 bytes a cada loop
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//Removendo o arquivo
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}