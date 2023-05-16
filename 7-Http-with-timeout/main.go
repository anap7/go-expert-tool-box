package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

/*
	http client do Go e suas chamadas

	Pensando em performance para chamada e respostas de
	um http, vamos trabalhar com os "limites" de chamadas ou
	o famoso timeout
*/
func main() {
	//Time de microsegundos e é esperado um panic
	c := http.Client{Timeout: time.Microsecond}
	res, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}