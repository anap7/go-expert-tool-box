package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			fmt.Println("estou na thread 02")
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println("estou na thread 01")
		fmt.Println(x)
	}
	
}