package main
import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) setNameWithoutPointer(name string) {
	p.Name = name
	fmt.Println("On method setName:", p.Name)
}

func (p *Person) setNameWithPointer(name string) {
	p.Name = name
	fmt.Println("On method setName:", p.Name)
}

func main() {
	println("-----------------------------------")
	println("Example without pointers")
	p := Person{
		Name: "Anaf",
		Age: 24,
	}
	
	p.setNameWithoutPointer("Vinicius")
	fmt.Println("On main function:", p.Name) //Expected: Ana

	println("-----------------------------------")
	println("Example with pointers")

	p2 := Person{
		Name: "Anaf",
		Age: 24,
	}
	
	p2.setNameWithPointer("Vinicius")
	fmt.Println("On main function:", p2.Name)
	println("-----------------------------------") //Expected: Vinicius

}

func simpleExampleWithPointers() {	
	name := "Ana"

	//Um ponteiro e indicando que essa variável é um apontamento na memoria
	var name2 *string 
	name2 = &name //Passando o endereço de memória para o ponteiro
	
	fmt.Println(&name2) //Mostra o endereço de memória
	fmt.Println(*name2) //Mostra o conteúdo que está nesse endereço de memória, isso se chama reference

	println("-----------------------------------------------------")

	/*Nesse caso ao atribuir um novo valor na reference do endereço de memória, significa
	que estou alterando o conteúdo dela. O endereço será o mesmo, porém o conteúdo será diferente*/
	*name2  = "Vinicius"
	fmt.Println(&name2)
	fmt.Println(*name2)
}