package main

import (
	"log"
	"net/http"
)

//Trabalhando com servido de arquivos, passando o arquivo como resposta http
func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Bloguinho!"))
	})
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}