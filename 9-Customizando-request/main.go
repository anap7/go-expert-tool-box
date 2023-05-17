package main

import (
	"io/ioutil"
	"net/http"
)

/*
	Customizando a requisição

	Vamos criar uma requisição e configurar a nossa request,
	criando uma request, passando para o cliente e executando ele com
	o método Do
*/
func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	/*Meu cliente recebendo a request no seu método Do e
	tudo que fizemos de configuração na request será configurada
	por ele*/
	res, err := c.Do(req)
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