package main

import (
	"log"
	"net/http"
	"time"
)

/*
	Objetivo
		Criar uma request cujo contexto tem o timeout de 5 segundos
*/

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//Pegando o contexto da nossa requisição
	ctx := r.Context()
	log.Println("Request Started!")
	defer log.Println("Request Finished!")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request was processed successfully")
		w.Write([]byte("Request was processed successfully"))
	case <-ctx.Done():
		log.Println("Request was canceled by client")
	}
}