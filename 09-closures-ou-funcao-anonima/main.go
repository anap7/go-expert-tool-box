package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 2, 5, 7, 954, 24, 3, 34, 1, 33, 934, 12)
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}