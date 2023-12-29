package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) walk() (string, error) {
	if p.Name != "Ana" {
		return "", errors.New("name should be Ana")
	}
	return p.Name + " walked", nil
}

func main() {
	p := Person{
		Name: "Anaf",
		Age: 24,
	}
	res, err := p.walk()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}