package main

func SomaInteiro(m map[string]int) int{
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64{
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type Number interface {
	int | float64
}

type MyNumber int 

//Criando um generic que aceita um int ou um float
func SomaGeneric[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaGenericUsingNumberInterface[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//O comparable é uma interface que implementa tipos de variáveis comparativas
func Comparar[T comparable] (a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"Wesley": 1000, "Joao": 2000, "Carlos": 3000}
	m2 := map[string]float64{"Wesley": 1000.45, "Joao": 2000.3, "Carlos": 3000.75}
	println(SomaInteiro(m1))
	println(SomaFloat(m2))

	println(SomaGeneric(m1))
	println(SomaGeneric(m2))
	println(Comparar(10, 40))
}