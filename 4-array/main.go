package main 

import "fmt"
//Declaração de constantes
const a = "Hello, World"

//Declaração da sua própria tipagem
type ID int

//Declarando valores iniciais
var (
	b2 bool = true
	c2 int = 10
	d2 string = "Wesley"
	e2 float64 = 1.2
	f2 ID = 1
)

func main() {
	//Normalmente o array tem valores fixos
	
	//Interagindo com um array
	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 10	

	fmt.Println(myArray[1]) //Pegando o valor do array
	fmt.Println(len(myArray)) //Pegando o total de posições do Array
	fmt.Println(len(myArray) - 1) //Pegando a última posição
	fmt.Println(myArray[len(myArray) - 1]) //Pegando o valor da última posição

	//Percorrendo um array
	var myArray2 [3]int

	myArray2[0] = 10
	myArray2[1] = 20
	myArray2[2] = 30

	for i, v := range myArray2 {
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}

}
