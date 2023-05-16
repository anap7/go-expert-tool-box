package main 

//Declaração de constantes
const a = "Hello, World"

//Declaração da sua própria tipagem
type ID int

/*Ao declarar variáveis com tipos o go já insere um valor inicial */
var (
	b bool //Por padrão insere "false"
	c int //Por padrão insere 0
	d string //Por padrão é vazio
	e float64 //Por padrão o valor é +0.000000e+000
	f ID = 1 //Declarando o próprio tipo baseado no "type" do go
)

//Declarando valores iniciais
var (
	b2 bool = true
	c2 int = 10
	d2 string = "Wesley"
	e2 float64 = 1.2
)

func main() {
	b = true
	//Declaração de uma variável e atribui - Ele é feito na primeira vez que a variável é criada
	a2 := "X"
	a2 = "oioi"
	println(a2)
	println(e)
	println(f)
}
