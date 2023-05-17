package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

/*
	Objetivo
		Um client para interagir com o nosso server
*/

func main() {
	//Iniciando contexto com timeout de 3 segundos de espera
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
