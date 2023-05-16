package main

import (
	"html/template"
	"net/http"
)

/*
	Introdução

	Em algum momento você já construiu uma aplicação dinâmica? Exemplo, precisar
	gerar um HTML dinâmico com os valores do banco de dados ou um e-mail com um
	código dinâmico

	O Golang possui um sistema de templates embutidos feitos para isso. Iremos fazer vários
	exemplos de templates, nesse arquivo iremos trabalhar com um arquivo HTML e listar
	dinamicamente os cursos via http request: Acesse http://localhost:8282/template-show
*/

type Course struct {
	Name string
	Workload int
}

type Courses []Course

func main() {
	http.HandleFunc("/template-show", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))

		//No exemplo anterior usamos o Execute para o os.Stdout e agora vamos jogar no nosso response do http
		err := tmp.Execute(w, Courses{
			{"Golang forever", 40},
			{"Cursos da Web para Jr JavaScript Embarcar na área de TI em 6 meses ganhando 100 mil por minuto", 10},
			{"Curso de Reacto para ganhar muito dinheiro pq não tem ninguem que sabe esse framework", 40},
			{"PHP", 10},
			{"Javao Legado", 40},
			{"C# nao eh c# nao viu", 80},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)
}