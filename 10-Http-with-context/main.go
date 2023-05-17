package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

/*
	http client do Go e suas chamadas

	Pensando em performance para chamada e respostas de
	um http, vamos trabalhar com os "limites" de chamadas ou
	o famoso timeout
*/
func main() {
	c := http.Client{}
	/*É necessario "bufferizar" um slice de bytes para que seja permitido passar
	como body no método Post, pois o body é um io.Reader*/
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Milta"}`))
	/*É esperado o status 405 Not Allowed*/
	res, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}