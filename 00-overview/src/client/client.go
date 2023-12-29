package main

import (
	"fmt"
	"time"
)

func main() {
	requestId := make(chan int)
	concurrency := 2
	for i := 1; i <= concurrency; i++ {
		println("Dentro do Worker...")
		go worker(requestId, i)
	}
	for i := 0; i < 20; i++ {
		println("Mudando o request ID ....")
		requestId <- i
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {
		/*res, err := http.Get("http://localhost:8585/product")
		if err != nil {
			log.Fatal("...Error")
		}
		defer res.Body.Close()

		content, _ := io.ReadAll(res.Body)*/
		//fmt.Printf("Worker: %d. RequestId: %d. Content: %s", worker, r, string(content))
		fmt.Printf("Worker: %d. RequestId: %d \n", worker, r)
		time.Sleep(time.Second * 2)
	}
}