package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	//Opcao01
	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!!"))
	})*/

	var b blog
	b.title = "Narutinho" 

	mux.HandleFunc("/handler", HomeHandler)
	mux.HandleFunc("/server", b.ServerHttp)

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's me Mario!"))
}

type blog struct{
	title string
}

func (b blog) ServerHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}