package main

import "fmt"

func main() {
	a := 80
	t := 6
	r := float64(t) / float64(a)
	fmt.Println(r)


	s := []int{10, 20, 30, 50, 60, 70, 80, 90, 100}
	fmt.Println(s[4])
	s[4] = 600
	fmt.Println(s[4])

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)//Tamanho, capacidade e valores

	// Pega os valores da esquerda para frente, ou seja, retorna sem valor com o tamanho 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) 

	// Pega os valores da esquerda para frente, ou seja, pega os 4 primeiros valores que iniciam pela esquerda
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Ignora os primeiros 2 valores que iniciam pela esquerda e mostra os valores para frente
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(s[5:]), cap(s[5:]), s[5:])

	s[5] = 700

	for i := 0; i < len(s); i++ {
		fmt.Print("\n")
		fmt.Println(s[i])
	}

	for index, value := range s {
		fmt.Print("\n")
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	meuSlice := make([]int, 0)     // Slice tamanho 0
	meuSlice2 := make([]int, 0, 4)  // Slice tamanho 0 e capacidade 4

	meuSlice = append(meuSlice2, 500, 22, 41, 15, 35, 60, 85, 90, 110, 115)
	meuSlice[2] = 300
	fmt.Print("\n")
	fmt.Println(meuSlice)

	meuSlice2 = append(meuSlice2, 300, 20, 40, 15)
	fmt.Print("\n")
	fmt.Println(meuSlice2)
}