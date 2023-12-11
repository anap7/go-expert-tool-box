package main

import (
	"fmt"
	"time"
)

func counter(typ string) {
	for i := 0; i < 5; i++ {
		fmt.Println(typ, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go counter("worker 1 ->")
	go counter("worker 2 ->")
	time.Sleep(time.Second * 10)
}

func test1() {
	counter("without go routine")
	// Cria uma thread ao rodar
	go counter("with go routine")

	fmt.Println("Hello 1")
	fmt.Println("Hello 2")

	time.Sleep(time.Second)
	fmt.Println("stopped...")
}