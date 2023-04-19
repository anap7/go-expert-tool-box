package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	//Struct para JSON
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	//JSON para Struct
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaX)
	println(contaX.Saldo)
}