package main

import (
	"io"
	"net/http"
)

func main() {
	request, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}	
	println(string(res))
	request.Body.Close()
}