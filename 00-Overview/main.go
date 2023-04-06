package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	user := Person{
		Name: "Dirce",
		Age: 30,
	}
	fmt.Println(user.Name)
}