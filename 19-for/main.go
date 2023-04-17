package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"um", "dois", "tres"}

	for _, v := range numbers {
		println(v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
